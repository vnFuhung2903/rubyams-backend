package model

type Class struct {
	Id         uint       `gorm:"primaryKey"`
	Course     string     `gorm:"not null;index"`
	Teacher    string     `gorm:"not null;index"`
	Semester   uint       `gorm:"unique"`
	Students   []*Student `gorm:"many2many:students_classes"`
	Capacity   uint
	Registered uint
	Ended      uint
}
