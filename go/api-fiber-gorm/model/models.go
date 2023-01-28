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
	Department string  `gorm:"uniqueIndex; not null" json:"department"`
	Country    Country `gorm:"foreignKey:CountryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	CountryId  uint
}

// City struct
type City struct {
	gorm.Model
	City         string     `gorm:"uniqueIndex; not null" json:"city"`
	Department   Department `gorm:"foreignKey:DepartmentId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	DepartmentId uint
}

// Address struct
type Address struct {
	gorm.Model
	Street       string  `gorm:"not null" json:"street"`
	Description  string  `json:"description"`
	Country      Country `gorm:"foreignKey:CountryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	CountryId    uint
	Department   Department `gorm:"foreignKey:DepartmentId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	DepartmentId uint
	City         City `gorm:"foreignKey:CityId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	CityId       uint
	ZipCode      string `json:"zip_code"`
	Phone        string `json:"phone"`
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
	FirstName  string   `json:"fistname"`
	LastName   string   `json:"lastname"`
	UserType   UserType `gorm:"foreignKey:UserTypeId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; default:1"`
	UserTypeId uint
}

// UserAdress struct
type UserAddress struct {
	gorm.Model
	User      User `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	UserId    uint
	Address   Address `gorm:"foreignKey:AddressId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	AddressId uint
}

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

// UserPet struct // transfer pet
type UserPet struct {
	gorm.Model
	User   User `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	UserId uint
	Pet    Pet `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	PetId  uint
	State  bool `gorm:"default:true" json:"state"`
	// Visited_at time.Time `json:"registered"`
}

// User struct
type PetVaccine struct {
	gorm.Model
	Date       time.Time `json:"date"`
	Vaccine    string    `gorm:"not null" json:"vaccine"`
	Vaccinator string    `json:"vaccinator"`
	Pet        Pet       `gorm:"foreignKey:PetId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	PetId      uint
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
