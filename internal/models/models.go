package models

import (
	"time"
)

type UserType struct {
	ID   int    `json:"id" gorm:"primaryKey;column:usertype"`
	Name string `json:"name" gorm:"column:usertype_name"`
}

func (UserType) TableName() string {
	return "usertype"
}

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;column:cd_user"`
	Email        string    `json:"email" gorm:"unique;not null;size:320"`
	Password     string    `json:"-" gorm:"not null;size:60"`
	UserStatus   string    `json:"user_status" gorm:"size:24;default:'Registered'"`
	UserType     int       `json:"user_type" gorm:"column:ty_user"`
	UserSubtype  int       `json:"user_subtype" gorm:"column:subtype_user;default:0"`
	FirstName    string    `json:"first_name" gorm:"size:64;default:''"`
	LastName     string    `json:"last_name" gorm:"size:64;default:''"`
	PhoneNumber  string    `json:"phone_number" gorm:"size:30;default:''"`
	Gender       string    `json:"gender" gorm:"size:16;default:'Other'"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"default:'1900-01-01'"`
	WeChatID     string    `json:"wechat_id" gorm:"size:64;default:''"`
	Languages    string    `json:"languages" gorm:"size:128;default:''"`
	Occupation   string    `json:"occupation" gorm:"size:128;default:''"`
	Religion     string    `json:"religion" gorm:"size:64;default:''"`
	HeightCm     int       `json:"height_cm" gorm:"default:0"`
	WeightKg     int       `json:"weight_kg" gorm:"default:0"`
	OTPCode      string    `json:"-" gorm:"size:6;default:''"`
	OTPCreatedAt time.Time `json:"-" gorm:"default:'1970-01-01 00:00:01'"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

type Assessment struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:cd_assessment"`
	UserID    uint      `json:"user_id" gorm:"column:cd_user"`
	Status    int       `json:"status" gorm:"default:0"`
	DiseaseID int       `json:"disease_id" gorm:"column:cd_disease"`
	ProductID int       `json:"product_id" gorm:"column:cd_product"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type Appointment struct {
	ID              uint      `json:"id" gorm:"primaryKey;column:cd_appointment"`
	DoctorID        uint      `json:"doctor_id" gorm:"column:cd_doctor"`
	PatientID       uint      `json:"patient_id" gorm:"column:cd_user"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime string    `json:"appointment_time" gorm:"type:time"`
	DurationMinutes int       `json:"duration_minutes" gorm:"default:30"`
	Status          string    `json:"status" gorm:"size:32;default:'scheduled'"`
	Notes           string    `json:"notes" gorm:"type:text;default:''"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Doctor          User      `json:"doctor,omitempty" gorm:"foreignKey:DoctorID"`
	Patient         User      `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}

type Order struct {
	ID              uint      `json:"id" gorm:"primaryKey;column:cd_order"`
	UserID          uint      `json:"user_id" gorm:"column:cd_user"`
	TotalAmount     float64   `json:"total_amount" gorm:"type:decimal(10,2)"`
	Status          string    `json:"status" gorm:"size:32;default:'pending'"`
	OrderReference  string    `json:"order_reference" gorm:"size:64;default:''"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User            User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (Order) TableName() string {
	return "\"order\""
}

type ChatRoom struct {
	ID          uint      `json:"id" gorm:"primaryKey;column:cd_chat_room"`
	PatientID   uint      `json:"patient_id" gorm:"column:cd_patient"`
	StaffID     *uint     `json:"staff_id" gorm:"column:cd_staff"`
	RoomType    int       `json:"room_type" gorm:"column:cd_room_type"`
	Status      string    `json:"status" gorm:"size:32;default:'waiting'"`
	Subject     string    `json:"subject" gorm:"size:256;default:''"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Patient     User      `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Staff       *User     `json:"staff,omitempty" gorm:"foreignKey:StaffID"`
}

type ChatMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:cd_message"`
	RoomID    uint      `json:"room_id" gorm:"column:cd_chat_room"`
	UserID    uint      `json:"user_id" gorm:"column:cd_user"`
	Content   string    `json:"content" gorm:"type:text"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
