package entity

import (
	"time"
)

type Student struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"unique"`
	Name          string    `gorm:"not null"`
	StudentNumber uint      `gorm:"not null"`
	NFTAddress    string    `gorm:"unique"`
	AcademicYear  uint      `gorm:"not null; index"`
	Courses       []*Course `gorm:"many2many:students_courses"`
	EnrolledAt    time.Time `gorm:"not null"`
	DeletedAt     time.Time
	Wallet        string
	Graduated     bool
}
