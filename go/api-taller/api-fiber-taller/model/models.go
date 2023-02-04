package model

import (
	"time"

	"gorm.io/gorm"
)

type Estudiante struct {
	gorm.Model
	Nombre      string    `json:"nombre"`
	Apellido    string    `json:"apellido"`
	Edad        int       `json:"edad"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	About       string    `json:"about"`
	Matriculado time.Time `json:"matriculado"`
}

type Curso struct {
	gorm.Model
	Curso string `gorm:"not null" json:"curso"`
}

type EstudianteCurso struct {
	gorm.Model
	EstudianteId int        `json:"estudiante_id"`
	CursoId      int        `json:"curso_id"`
	Nota         float64    `json:"nota"`
	Estudiante   Estudiante `gorm:"foreignKey:EstudianteId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	Curso        Curso      `gorm:"foreignKey:CursoId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}
