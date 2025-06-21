package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed frontend/dist/*
var embedFrontend embed.FS

// Database Models
type UserType struct {
	UserType     int    `json:"userType" gorm:"primaryKey;column:user_type"`
	UserTypeName string `json:"userTypeName" gorm:"column:user_type_name"`
}

func (UserType) TableName() string { return "usertype" }

type User struct {
	CdUser        uint       `json:"cdUser" gorm:"primaryKey;column:cd_user"`
	Email         string     `json:"email" gorm:"column:email;unique;not null"`
	Password      string     `json:"-" gorm:"column:password;not null"`
	FirstName     string     `json:"firstName" gorm:"column:first_name"`
	LastName      string     `json:"lastName" gorm:"column:last_name"`
	UserStatus    string     `json:"userStatus" gorm:"column:user_status;default:'Pending'"`
	TyUser        int        `json:"tyUser" gorm:"column:ty_user;default:1"`
	PhoneNumber   string     `json:"phoneNumber" gorm:"column:phone_number"`
	Gender        string     `json:"gender" gorm:"column:gender;default:'Other'"`
	DateOfBirth   time.Time  `json:"dateOfBirth" gorm:"column:date_of_birth"`
	OtpCode       string     `json:"-" gorm:"column:otp_code"`
	OtpCreatedAt  time.Time  `json:"otpCreatedAt" gorm:"column:otp_created_at"`
	CdCountry     int        `json:"cdCountry" gorm:"column:cd_country"`
	CdState       int        `json:"cdState" gorm:"column:cd_state"`
	CdCity        int        `json:"cdCity" gorm:"column:cd_city"`
	StreetAddress string     `json:"streetAddress" gorm:"column:street_address"`
	PostalCode    string     `json:"postalCode" gorm:"column:postal_code"`
	ProfileComplete bool     `json:"profileComplete" gorm:"column:profile_complete;default:false"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

func (User) TableName() string { return "users" }

type Country struct {
	CdCountry   int    `json:"cdCountry" gorm:"primaryKey;column:cd_country"`
	CountryName string `json:"countryName" gorm:"column:country_name"`
	CountryAbbr string `json:"countryAbbr" gorm:"column:country_abbr"`
}

func (Country) TableName() string { return "country" }

type State struct {
	CdCountry int    `json:"cdCountry" gorm:"primaryKey;column:cd_country"`
	CdState   int    `json:"cdState" gorm:"primaryKey;column:cd_state"`
	StateName string `json:"stateName" gorm:"column:state_name"`
	StateAbbr string `json:"stateAbbr" gorm:"column:state_abbr"`
}

func (State) TableName() string { return "state" }

type City struct {
	CdCountry int    `json:"cdCountry" gorm:"primaryKey;column:cd_country"`
	CdState   int    `json:"cdState" gorm:"primaryKey;column:cd_state"`
	CdCity    int    `json:"cdCity" gorm:"primaryKey;column:cd_city"`
	CityName  string `json:"cityName" gorm:"column:city_name"`
	CityAbbr  string `json:"cityAbbr" gorm:"column:city_abbr"`
}

func (City) TableName() string { return "city" }

var db *gorm.DB

func main() {
	// Initialize database
	initDB()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Printf("Error: %v", err)
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} ${path}\n",
	}))
	
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		dbStatus := "disconnected"
		userCount := int64(0)
		
		if db != nil {
			sqlDB, err := db.DB()
			if err == nil && sqlDB.Ping() == nil {
				dbStatus = "connected"
				db.Model(&User{}).Count(&userCount)
			}
		}

		return c.JSON(fiber.Map{
			"status":     "ok",
			"message":    "VCM Medical Platform API",
			"version":    "2.0.0",
			"database":   dbStatus,
			"userCount":  userCount,
			"time":       time.Now(),
			"features": []string{
				"User Registration",
				"Email OTP Verification",
				"JWT Authentication",
				"Profile Management",
				"Location Data",
				"React Frontend",
			},
		})
	})

	// Setup API routes
	setupRoutes(app)

	// Setup frontend serving - this is the key fix
	setupFrontend(app)

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("🚀 VCM Medical Platform v2.0 starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func setupFrontend(app *fiber.App) {
	// Try to serve embedded frontend
	frontendFS, err := fs.Sub(embedFrontend, "frontend/dist")
	if err != nil {
		log.Printf("❌ Failed to embed frontend: %v", err)
		// Fallback to filesystem serving
		setupFilesystemFrontend(app)
		return
	}

	// Serve static files from embedded filesystem
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendFS),
		PathPrefix: "assets",
		Browse:     false,
	}))

	// Serve index.html for all routes (SPA support)
	app.Get("/*", func(c *fiber.Ctx) error {
		path := c.Path()
		
		// Skip API routes
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/health") {
			return c.Next()
		}

		// Try to serve the requested file
		if strings.Contains(path, ".") {
			file, err := frontendFS.Open(strings.TrimPrefix(path, "/"))
			if err == nil {
				defer file.Close()
				return c.SendStream(file)
			}
		}

		// Serve index.html for all other routes (SPA routing)
		indexFile, err := frontendFS.Open("index.html")
		if err != nil {
			log.Printf("❌ Failed to open index.html: %v", err)
			return c.Status(404).SendString("Frontend not found")
		}
		defer indexFile.Close()

		c.Set("Content-Type", "text/html")
		return c.SendStream(indexFile)
	})

	log.Println("✅ Serving embedded React frontend")
}

func setupFilesystemFrontend(app *fiber.App) {
	// Fallback to filesystem if embed fails
	distPath := "./frontend/dist"
	
	// Check if dist directory exists
	if _, err := os.Stat(distPath); os.IsNotExist(err) {
		log.Printf("❌ Frontend dist directory not found: %s", distPath)
		// Serve fallback HTML
		app.Get("/*", func(c *fiber.Ctx) error {
			if strings.HasPrefix(c.Path(), "/api") || strings.HasPrefix(c.Path(), "/health") {
				return c.Next()
			}
			return c.Type("html").SendString(getFallbackHTML())
		})
		return
	}

	// Serve static files
	app.Static("/", distPath)

	// Serve index.html for all other routes (SPA routing)
	app.Get("/*", func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") || strings.HasPrefix(c.Path(), "/health") {
			return c.Next()
		}
		return c.SendFile(distPath + "/index.html")
	})

	log.Println("✅ Serving React frontend from filesystem")
}

func setupRoutes(app *fiber.App) {
	// API routes
	api := app.Group("/api/v1")
	
	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", register)
	auth.Post("/verify-otp", verifyOTP)
	auth.Post("/resend-otp", resendOTP)
	auth.Post("/login", login)
	auth.Post("/complete-profile", completeProfile)
	
	// Protected auth routes
	protected := auth.Group("", jwtMiddleware)
	protected.Get("/me", getProfile)
	protected.Put("/profile", updateProfile)

	// Location routes
	location := api.Group("/location")
	location.Get("/countries", getCountries)
	location.Get("/states/:countryId", getStates)
	location.Get("/cities/:countryId/:stateId", getCities)
}

func initDB() {
	var err error
	
	// Get database URL (Railway provides this automatically)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("⚠️ No DATABASE_URL found - using fallback connection")
		// Fallback for local development
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "vcm_user"), 
			getEnv("DB_PASSWORD", "vcm_password_2024"),
			getEnv("DB_NAME", "vcm_medical"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_SSLMODE", "disable"),
		)
	}

	// Connect to database
	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Database connection failed: %v", err)
		log.Println("🔄 Continuing without database...")
		return
	}

	log.Println("✅ Database connected successfully!")

	// Auto migrate tables
	err = db.AutoMigrate(
		&UserType{},
		&User{},
		&Country{},
		&State{},
		&City{},
	)
	if err != nil {
		log.Printf("⚠️ Migration warning: %v", err)
	} else {
		log.Println("✅ Database migration completed")
	}

	// Seed initial data
	seedData()
}

// Auth Handlers
func register(c *fiber.Ctx) error {
	var req struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		UserType    int    `json:"userType"`
		PhoneNumber string `json:"phoneNumber"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"dateOfBirth"`
		CountryId   int    `json:"countryId"`
		StateId     int    `json:"stateId"`
		CityId      int    `json:"cityId"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email, password, first name, and last name are required"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	// Check if user already exists
	var existingUser User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already registered"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Generate OTP
	otp := generateOTP()

	// Parse date of birth
	var dob time.Time
	if req.DateOfBirth != "" {
		dob, _ = time.Parse("2006-01-02", req.DateOfBirth)
	}

	// Create user
	user := User{
		Email:         strings.ToLower(req.Email),
		Password:      string(hashedPassword),
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		TyUser:        req.UserType,
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
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// In production, send email here
	log.Printf("📧 OTP for %s: %s (expires in 10 minutes)", req.Email, otp)

	return c.JSON(fiber.Map{
		"message": "Registration successful! Please check your email for OTP verification.",
		"email":   req.Email,
		"userId":  user.CdUser,
		// Remove this in production:
		"otp": otp,
	})
}

func verifyOTP(c *fiber.Ctx) error {
	var req struct {
		Email   string `json:"email"`
		OtpCode string `json:"otpCode"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.Where("email = ? AND otp_code = ?", strings.ToLower(req.Email), req.OtpCode).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid OTP or email"})
	}

	// Check if OTP is expired (10 minutes)
	if time.Since(user.OtpCreatedAt) > 10*time.Minute {
		return c.Status(400).JSON(fiber.Map{"error": "OTP has expired. Please request a new one."})
	}

	// Update user status
	user.UserStatus = "Verified"
	user.OtpCode = ""
	
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to verify user"})
	}

	// Generate JWT token
	token, err := generateJWT(user.CdUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"message": "Email verified successfully!",
		"token":   token,
		"user":    user,
		"requiresProfileCompletion": !user.ProfileComplete,
	})
}

func login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.Where("email = ?", strings.ToLower(req.Email)).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Check if email is verified
	if user.UserStatus == "Pending" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Please verify your email first",
			"needsVerification": true,
			"email": user.Email,
		})
	}

	// Generate JWT token
	token, err := generateJWT(user.CdUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful!",
		"token":   token,
		"user":    user,
		"requiresProfileCompletion": !user.ProfileComplete,
	})
}

func completeProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	
	var req struct {
		PhoneNumber   string `json:"phoneNumber"`
		Gender        string `json:"gender"`
		DateOfBirth   string `json:"dateOfBirth"`
		CountryId     int    `json:"countryId"`
		StateId       int    `json:"stateId"`
		CityId        int    `json:"cityId"`
		StreetAddress string `json:"streetAddress"`
		PostalCode    string `json:"postalCode"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Parse date of birth
	if req.DateOfBirth != "" {
		if dob, err := time.Parse("2006-01-02", req.DateOfBirth); err == nil {
			user.DateOfBirth = dob
		}
	}

	// Update user profile
	user.PhoneNumber = req.PhoneNumber
	user.Gender = req.Gender
	user.CdCountry = req.CountryId
	user.CdState = req.StateId
	user.CdCity = req.CityId
	user.StreetAddress = req.StreetAddress
	user.PostalCode = req.PostalCode
	user.ProfileComplete = true
	user.UserStatus = "Active"

	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{
		"message": "Profile completed successfully!",
		"user":    user,
	})
}

func resendOTP(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.Where("email = ?", strings.ToLower(req.Email)).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Email not found"})
	}

	if user.UserStatus != "Pending" {
		return c.Status(400).JSON(fiber.Map{"error": "Email already verified"})
	}

	// Generate new OTP
	otp := generateOTP()
	user.OtpCode = otp
	user.OtpCreatedAt = time.Now()

	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to resend OTP"})
	}

	log.Printf("📧 New OTP for %s: %s", req.Email, otp)

	return c.JSON(fiber.Map{
		"message": "New OTP sent to your email",
		"otp":     otp, // Remove in production
	})
}

func getProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	
	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func updateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if db == nil {
		return c.Status(503).JSON(fiber.Map{"error": "Database not available"})
	}

	if err := db.Model(&User{}).Where("cd_user = ?", userID).Updates(updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

// Location Handlers
func getCountries(c *fiber.Ctx) error {
	if db == nil {
		// Return static data if database not available
		return c.JSON(fiber.Map{
			"countries": getStaticCountries(),
		})
	}

	var countries []Country
	if err := db.Find(&countries).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch countries"})
	}

	return c.JSON(fiber.Map{"countries": countries})
}

func getStates(c *fiber.Ctx) error {
	countryId, err := strconv.Atoi(c.Params("countryId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid country ID"})
	}

	if db == nil {
		return c.JSON(fiber.Map{"states": []State{}})
	}

	var states []State
	if err := db.Where("cd_country = ?", countryId).Find(&states).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch states"})
	}

	return c.JSON(fiber.Map{"states": states})
}

func getCities(c *fiber.Ctx) error {
	countryId, _ := strconv.Atoi(c.Params("countryId"))
	stateId, _ := strconv.Atoi(c.Params("stateId"))

	if db == nil {
		return c.JSON(fiber.Map{"cities": []City{}})
	}

	var cities []City
	if err := db.Where("cd_country = ? AND cd_state = ?", countryId, stateId).Find(&cities).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch cities"})
	}

	return c.JSON(fiber.Map{"cities": cities})
}

// Utility Functions
func generateOTP() string {
	// Generate 6-digit OTP
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := getEnv("JWT_SECRET", "vcm-medical-platform-secret-key-2024")
	return token.SignedString([]byte(secret))
}

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header format"})
	}

	tokenString := authHeader[7:]
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := getEnv("JWT_SECRET", "vcm-medical-platform-secret-key-2024")
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	userID := uint(claims["user_id"].(float64))
	c.Locals("userID", userID)

	return c.Next()
}

func seedData() {
	if db == nil {
		return
	}

	// Seed user types
	userTypes := []UserType{
		{UserType: 1, UserTypeName: "Patient"},
		{UserType: 5, UserTypeName: "Doctor"},
		{UserType: 10, UserTypeName: "Nurse"},
		{UserType: 15, UserTypeName: "Administrator"},
	}

	for _, userType := range userTypes {
		db.FirstOrCreate(&userType, UserType{UserType: userType.UserType})
	}

	// Seed countries
	countries := []Country{
		{CdCountry: 1, CountryName: "United States", CountryAbbr: "US"},
		{CdCountry: 2, CountryName: "Canada", CountryAbbr: "CA"},
		{CdCountry: 3, CountryName: "United Kingdom", CountryAbbr: "UK"},
		{CdCountry: 4, CountryName: "Germany", CountryAbbr: "DE"},
		{CdCountry: 5, CountryName: "France", CountryAbbr: "FR"},
		{CdCountry: 6, CountryName: "Japan", CountryAbbr: "JP"},
		{CdCountry: 7, CountryName: "Australia", CountryAbbr: "AU"},
		{CdCountry: 8, CountryName: "India", CountryAbbr: "IN"},
	}

	for _, country := range countries {
		db.FirstOrCreate(&country, Country{CdCountry: country.CdCountry})
	}

	// Seed some sample states
	states := []State{
		{CdCountry: 1, CdState: 1, StateName: "California", StateAbbr: "CA"},
		{CdCountry: 1, CdState: 2, StateName: "New York", StateAbbr: "NY"},
		{CdCountry: 1, CdState: 3, StateName: "Texas", StateAbbr: "TX"},
		{CdCountry: 2, CdState: 1, StateName: "Ontario", StateAbbr: "ON"},
		{CdCountry: 2, CdState: 2, StateName: "Quebec", StateAbbr: "QC"},
		{CdCountry: 3, CdState: 1, StateName: "England", StateAbbr: "ENG"},
		{CdCountry: 3, CdState: 2, StateName: "Scotland", StateAbbr: "SCT"},
	}

	for _, state := range states {
		db.FirstOrCreate(&state, State{CdCountry: state.CdCountry, CdState: state.CdState})
	}

	log.Println("✅ Database seeding completed")
}

func getStaticCountries() []fiber.Map {
	return []fiber.Map{
		{"cdCountry": 1, "countryName": "United States", "countryAbbr": "US"},
		{"cdCountry": 2, "countryName": "Canada", "countryAbbr": "CA"},
		{"cdCountry": 3, "countryName": "United Kingdom", "countryAbbr": "UK"},
		{"cdCountry": 4, "countryName": "Germany", "countryAbbr": "DE"},
		{"cdCountry": 5, "countryName": "France", "countryAbbr": "FR"},
		{"cdCountry": 6, "countryName": "Japan", "countryAbbr": "JP"},
		{"cdCountry": 7, "countryName": "Australia", "countryAbbr": "AU"},
		{"cdCountry": 8, "countryName": "India", "countryAbbr": "IN"},
	}
}

func getFallbackHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            text-align: center;
            background: rgba(255,255,255,0.1);
            padding: 2rem;
            border-radius: 20px;
            backdrop-filter: blur(10px);
            max-width: 500px;
            margin: 1rem;
        }
        h1 { font-size: 2.5rem; margin-bottom: 1rem; }
        .success { 
            background: rgba(76, 175, 80, 0.3);
            padding: 1rem 2rem;
            border-radius: 25px;
            margin: 1rem 0;
            border: 2px solid rgba(76, 175, 80, 0.5);
        }
        .endpoint {
            background: rgba(0,0,0,0.2);
            padding: 0.5rem 1rem;
            margin: 0.5rem 0;
            border-radius: 8px;
            font-family: monospace;
        }
        a { color: #87CEEB; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🏥 VCM Medical Platform</h1>
        <p>Advanced Medical Treatment Platform</p>
        
        <div class="success">
            ⚠️ React Frontend Loading Issue
        </div>
        
        <div style="background: rgba(255,255,255,0.1); padding: 1rem; border-radius: 10px; margin: 1rem 0;">
            <h3>🔗 Test Endpoints:</h3>
            <div class="endpoint"><a href="/health">/health</a> - API Health Check</div>
            <div class="endpoint"><a href="/api/v1/location/countries">/api/v1/location/countries</a> - Countries Data</div>
        </div>
        
        <p style="font-size: 0.9rem; opacity: 0.8; margin-top: 2rem;">
            Frontend build issue detected. Deploying fix...<br>
            Backend API is working correctly.
        </p>
    </div>
</body>
</html>`
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
