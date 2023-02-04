package model

import (
	"time"

	"gorm.io/gorm"
)

// Country struct
type Country struct {
	gorm.Model
	Country string `gorm:"uniqueIndex; not null" json:"country"`
}

// Department struct
type Department struct {
	gorm.Model
	Department string  `gorm:"not null" json:"department"`
	CountryId  int     `json:"country_id"`
	Country    Country `gorm:"foreignKey:CountryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}

// City struct
type City struct {
	gorm.Model
	City         string     `gorm:"not null" json:"city"`
	DepartmentId int        `json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}

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
type UserAddress struct {
	gorm.Model
	UserId       int    `json:"user_id"`
	Name         string `gorm:"not null" json:"name"`
	CityId       int    `json:"city_id"`
	ZipCode      string `json:"zip_code"`
	Details      string `gorm:"not null" json:"details"`
	OtherDetails string `json:"other_details"`
	Phone        string `json:"phone"`
	Preferred    bool   `gorm:"default:false" json:"preferred"`
	City         City   `gorm:"foreignKey:CityId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	User         User   `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}

// UserAdress struct
// type UserAddress struct {
// 	gorm.Model
// 	UserId    int
// 	AddressId int
// 	Preferred bool    `gorm:"default:false" json:"preferred"`
// 	User      User    `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// 	Address   Address `gorm:"foreignKey:AddressId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
// }

// Pet struct
type Pet struct {
	gorm.Model
	Name      string    `gorm:"not null" json:"name"`
	Type      string    `gorm:"not null" json:"type"`
	Breed     string    `json:"breed"`
	Sex       string    `json:"sex"`
	Color     string    `json:"lastname"`
	BirthDate time.Time `json:"birth_date"`
}

// Onw struct // transfer pet
type Own struct {
	gorm.Model
	UserId int  `json:"user_id"`
	PetId  int  `json:"pet_id"`
	State  bool `gorm:"default:true" json:"state"`
	User   User `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	Pet    Pet  `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}

// PetRecord struct
type PetRecord struct {
	gorm.Model
	Date        time.Time `json:"date"`
	Record      string    `gorm:"not null" json:"record"`
	Description string    `json:"description"`
	PetId       int       `json:"pet_id"`
	Pet         Pet       `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}

// User struct
// type Vaccine struct {
// 	gorm.Model
// 	Username  string `gorm:"uniqueIndex; not null" json:"username"`
// 	Email     string `gorm:"uniqueIndex; not null" json:"email"`
// 	Password  string `gorm:"not null" json:"password"`
// 	FirstName string `json:"fistname"`
// 	LastName  string `json:"lastname"`
// }
