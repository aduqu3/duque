package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Cod        uint8  `gorm:"primaryKey"`
	Name       string `gorm: "not null"`
	SecondName string
	LastName   string `gorm:"not null; default: 'NN'"`
	Visited_at time.Time
}

// func confgDB(){
// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disabble TimeZone=America/Bogota"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	fmt.Println("DB conecction success")

// 	fmt.Println(db)
// 	// return db
// }

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=America/Bogota"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("DB conecction success")
	fmt.Println(db)

	// migrar modelo
	db.AutoMigrate(&Student{})

	// crear datos
	// db.Create(&Student{Cod: 1, Name: "Andres", LastName: "Duque"})
	// db.Create(&Student{Cod: 2, Name: "Alejandro", Visited_at: time.Now()})


	// eliminar todo deja logs
	// db.Where("1 = 1").Delete(&Student{})

	// eliminar rows de forma fea
	// db.Exec("DELETE from students")

	// eliminar todo deja logs
	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Student{})


	// consultas
	var stud1, stud2 Student
	// db.First(&stud1, 1)
	db.First(&stud1, "cod = ? ", 1)
	fmt.Println(stud1)

	db.First(&stud2, "cod = ? ", 4)
	fmt.Println(stud2)

	// update data
	db.Model(&stud1).Update("name", "Juan2")
	db.Model(&stud2).Updates(Student{Cod: 4, Name: "Margarita4"})

	// delte 
}
