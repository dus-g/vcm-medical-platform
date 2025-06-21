package models

import (
	"time"
	"gorm.io/gorm"
)

type UserType struct {
	UserType     int    `gorm:"primaryKey" json:"usertype"`
	UserTypeName string `gorm:"size:64;not null" json:"usertype_name"`
}

type User struct {
	CdUser       uint      `gorm:"primaryKey;autoIncrement" json:"cd_user"`
	UserStatus   string    `gorm:"size:24;not null;default:'Registered'" json:"user_status"`
	TyUser       int       `gorm:"not null" json:"ty_user"`
	SubtypeUser  int       `gorm:"not null;default:0" json:"subtype_user"`
	Email        string    `gorm:"size:320;uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"size:60;not null" json:"-"`
	
	// OTP Info
	OtpCode      string    `gorm:"size:6;not null;default:''" json:"-"`
	OtpCreatedAt time.Time `gorm:"default:'1970-01-01 00:00:01'" json:"-"`
	
	// Personal Information
	FirstName    string    `gorm:"size:64;not null;default:''" json:"first_name"`
	LastName     string    `gorm:"size:64;not null;default:''" json:"last_name"`
	Gender       string    `gorm:"size:16;not null;default:'Other'" json:"gender"`
	PhoneNumber  string    `gorm:"size:30;not null;default:''" json:"phone_number"`
	DateOfBirth  time.Time `gorm:"not null;default:'1900-01-01'" json:"date_of_birth"`
	WechatId     string    `gorm:"size:64;not null;default:''" json:"wechat_id"`
	
	Languages      string `gorm:"size:128;not null;default:''" json:"languages"`
	Occupation     string `gorm:"size:128;not null;default:''" json:"occupation"`
	Religion       string `gorm:"size:64;not null;default:''" json:"religion"`
	HeightCm       int    `gorm:"not null;default:0" json:"height_cm"`
	WeightKg       int    `gorm:"not null;default:0" json:"weight_kg"`
	MaritalStatus  string `gorm:"size:24;not null;default:'Single'" json:"marital_status"`
	NoChildren     int    `gorm:"not null;default:0" json:"no_children"`
	
	// Address
	CdCountry     int    `gorm:"not null;default:0" json:"cd_country"`
	CdState       int    `gorm:"not null;default:0" json:"cd_state"`
	CdCity        int    `gorm:"not null;default:0" json:"cd_city"`
	CdDistrict    int    `gorm:"not null;default:0" json:"cd_district"`
	CdStreet      int    `gorm:"not null;default:0" json:"cd_street"`
	StreetAddress string `gorm:"size:255;not null;default:''" json:"street_address"`
	PostalCode    string `gorm:"size:32;not null;default:''" json:"postal_code"`
	
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Associations
	UserType UserType `gorm:"foreignKey:TyUser;references:UserType" json:"user_type,omitempty"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

func (UserType) TableName() string {
	return "usertype"
}

// User methods
func (u *User) GetFullName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	return u.Email
}

func (u *User) IsProfileComplete() bool {
	return u.FirstName != "" && u.LastName != "" && u.PhoneNumber != ""
}

// Country and location models
type Country struct {
	CdCountry    int    `gorm:"primaryKey" json:"cd_country"`
	CountryName  string `gorm:"size:64;not null" json:"country_name"`
	CountryAbbr  string `gorm:"size:8;not null" json:"country_abbr"`
}

func (Country) TableName() string {
	return "country"
}

type State struct {
	CdCountry  int    `gorm:"primaryKey" json:"cd_country"`
	CdState    int    `gorm:"primaryKey" json:"cd_state"`
	StateName  string `gorm:"size:64;not null" json:"state_name"`
	StateAbbr  string `gorm:"size:16;not null" json:"state_abbr"`
}

func (State) TableName() string {
	return "state"
}

type City struct {
	CdCountry int    `gorm:"primaryKey" json:"cd_country"`
	CdState   int    `gorm:"primaryKey" json:"cd_state"`
	CdCity    int    `gorm:"primaryKey" json:"cd_city"`
	CityName  string `gorm:"size:64;not null" json:"city_name"`
	CityAbbr  string `gorm:"size:16;not null" json:"city_abbr"`
}

func (City) TableName() string {
	return "city"
}

type District struct {
	CdCountry     int    `gorm:"primaryKey" json:"cd_country"`
	CdState       int    `gorm:"primaryKey" json:"cd_state"`
	CdCity        int    `gorm:"primaryKey" json:"cd_city"`
	CdDistrict    int    `gorm:"primaryKey" json:"cd_district"`
	DistrictName  string `gorm:"size:64;not null" json:"district_name"`
	DistrictAbbr  string `gorm:"size:16;not null" json:"district_abbr"`
}

func (District) TableName() string {
	return "district"
}
