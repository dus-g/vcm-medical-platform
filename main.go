package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed frontend/dist
var embedFrontend embed.FS

// Database models matching your existing structure
type UserType struct {
	UserType     int    `json:"userType" gorm:"primaryKey;column:user_type"`
	UserTypeName string `json:"userTypeName" gorm:"column:user_type_name"`
}

func (UserType) TableName() string {
	return "usertype"
}

type User struct {
	CdUser        uint      `json:"cdUser" gorm:"primaryKey;column:cd_user"`
	UserStatus    string    `json:"userStatus" gorm:"column:user_status;default:'Registered'"`
	TyUser        int       `json:"tyUser" gorm:"column:ty_user"`
	SubtypeUser   int       `json:"subtypeUser" gorm:"column:subtype_user;default:0"`
	Email         string    `json:"email" gorm:"column:email;unique"`
	Password      string    `json:"-" gorm:"column:password"`
	OtpCode       string    `json:"-" gorm:"column:otp_code;default:''"`
	OtpCreatedAt  time.Time `json:"otpCreatedAt" gorm:"column:otp_created_at;default:'1970-01-01 00:00:01'"`
	FirstName     string    `json:"firstName" gorm:"column:first_name;default:''"`
	LastName      string    `json:"lastName" gorm:"column:last_name;default:''"`
	Gender        string    `json:"gender" gorm:"column:gender;default:'Other'"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"column:phone_number;default:''"`
	DateOfBirth   time.Time `json:"dateOfBirth" gorm:"column:date_of_birth;default:'1900-01-01'"`
	CdCountry     int       `json:"cdCountry" gorm:"column:cd_country;default:0"`
	CdState       int       `json:"cdState" gorm:"column:cd_state;default:0"`
	CdCity        int       `json:"cdCity" gorm:"column:cd_city;default:0"`
	StreetAddress string    `json:"streetAddress" gorm:"column:street_address;default:''"`
	PostalCode    string    `json:"postalCode" gorm:"column:postal_code;default:''"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}

type Country struct {
	CdCountry   int    `json:"cd_country" gorm:"primaryKey;column:cd_country"`
	CountryName string `json:"country_name" gorm:"column:country_name"`
	CountryAbbr string `json:"country_abbr" gorm:"column:country_abbr"`
}

func (Country) TableName() string {
	return "country"
}

type State struct {
	CdCountry int    `json:"cd_country" gorm:"primaryKey;column:cd_country"`
	CdState   int    `json:"cd_state" gorm:"primaryKey;column:cd_state"`
	StateName string `json:"state_name" gorm:"column:state_name"`
	StateAbbr string `json:"state_abbr" gorm:"column:state_abbr"`
}

func (State) TableName() string {
	return "state"
}

var db *gorm.DB

func main() {
	godotenv.Load()
	initDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		dbStatus := "disconnected"
		if db != nil {
			sqlDB, _ := db.DB()
			if sqlDB != nil && sqlDB.Ping() == nil {
				dbStatus = "connected"
			}
		}

		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "VCM Medical Platform API",
			"version":   "2.0.0",
			"database":  dbStatus,
			"time":      time.Now(),
		})
	})

	// API routes
	api := app.Group("/api/v1")
	setupRoutes(api)

	// Serve frontend
	setupFrontend(app)

	port := getEnv("PORT", "8080")
	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func initDB() {
	var err error
	
	// Use DATABASE_URL for Railway
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		// Fallback for local development
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "vcm_user"),
			getEnv("DB_PASSWORD", "vcm_password_2024"),
			getEnv("DB_NAME", "vcm_medical"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_SSLMODE", "disable"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Printf("❌ Database connection failed: %v", err)
		return
	}

	log.Println("✅ Database connected!")
}

func setupRoutes(api fiber.Router) {
	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", register)
	auth.Post("/login", login)
	auth.Post("/verify-otp", verifyOTP)
	auth.Post("/resend-otp", resendOTP)
	
	// Protected auth routes
	protected := auth.Group("", jwtMiddleware)
	protected.Get("/me", getProfile)

	// Location routes
	location := api.Group("/location")
	location.Get("/countries", getCountries)
	location.Get("/states/:countryId", getStates)
}

func setupFrontend(app *fiber.App) {
	// Try embedded frontend first
	frontendFS, err := fs.Sub(embedFrontend, "frontend/dist")
	if err == nil {
		app.Use("/", filesystem.New(filesystem.Config{
			Root:   http.FS(frontendFS),
			Browse: false,
		}))
		log.Println("✅ Serving embedded frontend")
		return
	}

	// Fallback to simple HTML
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
<!DOCTYPE html>
<html><head><title>VCM Medical Platform</title>
<style>body{font-family:Arial;padding:40px;background:linear-gradient(135deg,#667eea,#764ba2);color:white;text-align:center}</style>
</head><body>
<h1>🏥 VCM Medical Platform</h1>
<p>Advanced Medical Treatment Platform with 95% Efficacy</p>
<div style="background:rgba(76,175,80,0.2);padding:15px;border-radius:25px;margin:20px 0">✅ API Status: Running</div>
<div style="background:rgba(255,255,255,0.1);padding:20px;border-radius:15px">
<h3>🔗 API Endpoints:</h3>
<p>GET /health - Health Check</p>
<p>GET /api/v1/location/countries - Countries</p>
<p>POST /api/v1/auth/register - Registration</p>
<p>POST /api/v1/auth/login - Login</p>
</div>
</body></html>`)
	})
}

// Auth handlers
func register(c *fiber.Ctx) error {
	var req struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		UserType    int    `json:"userType"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"dateOfBirth"`
		CountryId   int    `json:"countryId"`
		StateId     int    `json:"stateId"`
		CityId      int    `json:"cityId"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Generate OTP
	otp := fmt.Sprintf("%06d", time.Now().Unix()%1000000)

	// Parse date of birth
	dob, _ := time.Parse("2006-01-02", req.DateOfBirth)

	user := User{
		Email:         req.Email,
		Password:      string(hashedPassword),
		TyUser:        req.UserType,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		PhoneNumber:   req.PhoneNumber,
		Gender:        req.Gender,
		DateOfBirth:   dob,
		CdCountry:     req.CountryId,
		CdState:       req.StateId,
		CdCity:        req.CityId,
		OtpCode:       otp,
		OtpCreatedAt:  time.Now(),
		UserStatus:    "Pending",
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Registration failed"})
	}

	log.Printf("📧 OTP for %s: %s", req.Email, otp)

	return c.JSON(fiber.Map{
		"message": "Registration successful",
		"email":   req.Email,
		"otp":     otp, // Remove in production
	})
}

func login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if user.UserStatus != "Active" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Please verify your email first",
			"needsVerification": true,
		})
	}

	token, err := generateJWT(user.CdUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func verifyOTP(c *fiber.Ctx) error {
	var req struct {
		Email   string `json:"email"`
		OtpCode string `json:"otpCode"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.Where("email = ? AND otp_code = ?", req.Email, req.OtpCode).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	if time.Since(user.OtpCreatedAt) > 10*time.Minute {
		return c.Status(400).JSON(fiber.Map{"error": "OTP expired"})
	}

	user.UserStatus = "Active"
	user.OtpCode = ""
	db.Save(&user)

	token, _ := generateJWT(user.CdUser)

	return c.JSON(fiber.Map{
		"message": "Email verified successfully",
		"token":   token,
		"user":    user,
	})
}

func resendOTP(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "OTP resent"})
}

func getProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func getCountries(c *fiber.Ctx) error {
	if db == nil {
		return c.JSON(fiber.Map{
			"countries": []fiber.Map{
				{"cd_country": 1, "country_name": "United States", "country_abbr": "US"},
				{"cd_country": 86, "country_name": "China", "country_abbr": "CN"},
			},
		})
	}

	var countries []Country
	db.Find(&countries)
	return c.JSON(fiber.Map{"countries": countries})
}

func getStates(c *fiber.Ctx) error {
	countryId, _ := strconv.Atoi(c.Params("countryId"))
	
	if db == nil {
		return c.JSON(fiber.Map{"states": []fiber.Map{}})
	}

	var states []State
	db.Where("cd_country = ?", countryId).Find(&states)
	return c.JSON(fiber.Map{"states": states})
}

// Utility functions
func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getEnv("JWT_SECRET", "default-secret")))
}

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header"})
	}

	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(getEnv("JWT_SECRET", "default-secret")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	c.Locals("userID", userID)

	return c.Next()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
