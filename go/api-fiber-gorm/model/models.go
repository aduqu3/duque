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
	Username   string   `gorm:"uniqueIndex; not null" json:"username"`
	Email      string   `gorm:"uniqueIndex; not null" json:"email"`
	Password   string   `gorm:"not null" json:"password"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	UserTypeId int      `gorm:"default:1" json:"user_type_id"`
	UserType   UserType `gorm:"foreignKey:UserTypeId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Address struct
// type UserAddress struct {
// 	gorm.Model
// 	UserId       int    `json:"user_id"`
// 	Name         string `gorm:"not null" json:"name"`
// 	Country      string `json:"country"`
// 	Region       string `json:"region"`
// 	City         string `json:"city"`
// 	ZipCode      string `json:"zip_code"`
// 	Details      string `gorm:"not null" json:"details"`
// 	OtherDetails string `json:"other_details"`
// 	Phone        string `json:"phone"`
// 	Preferred    bool   `gorm:"default:false" json:"preferred"`
// 	User         User   `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// }

// // UserAdress struct
// // type UserAddress struct {
// // 	gorm.Model
// // 	UserId    int
// // 	AddressId int
// // 	Preferred bool    `gorm:"default:false" json:"preferred"`
// // 	User      User    `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// // 	Address   Address `gorm:"foreignKey:AddressId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// // }

// // Pet struct
// type Pet struct {
// 	gorm.Model
// 	Name      string    `gorm:"not null" json:"name"`
// 	Type      string    `gorm:"not null" json:"type"`
// 	Breed     string    `json:"breed"`
// 	Gender    string    `json:"gender"`
// 	Color     string    `json:"color"`
// 	BirthDate time.Time `json:"birth_date"`
// 	Image     string    `json:"image"`
// }

// // Onw struct // transfer pet
// type Own struct {
// 	gorm.Model
// 	UserId int  `json:"user_id"`
// 	PetId  int  `json:"pet_id"`
// 	User   User `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// 	Pet    Pet  `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// }

// // PetRecord struct
// type PetRecord struct {
// 	gorm.Model
// 	Date        time.Time `json:"date"`
// 	Record      string    `gorm:"not null" json:"record"`
// 	Description string    `json:"description"`
// 	PetId       int       `json:"pet_id"`
// 	Pet         Pet       `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// }
