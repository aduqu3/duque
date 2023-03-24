package model

import (
	"gorm.io/gorm"
)

// UserType struct
type UserType struct {
	gorm.Model
	Role string `gorm:"uniqueIndex; not null" json:"role"`
}

// User struct
type User struct {
	gorm.Model
	Username   string   `gorm:"not null; uniqueIndex" json:"username"`
	Email      string   `gorm:"not null; uniqueIndex" json:"email"`
	Password   string   `gorm:"not null" json:"password"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	UserTypeId int      `gorm:"default:1" json:"user_type_id"`
	UserType   UserType `gorm:"foreignKey:UserTypeId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
