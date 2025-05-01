package entity

import "time"

type Course struct {
	ID          uint         `gorm:"primaryKey"`
	CourseName  string       `gorm:"not null;index"`
	TeacherName string       `gorm:"not null;index"`
	Semester    float32      `gorm:"index"`
	Students    []*Student   `gorm:"many2many:students_courses"`
	InProgress  bool         `gorm:"index"`
	Weekday     time.Weekday `gorm:"index"`
	StartTime   time.Time
	EndTime     time.Time
	Capacity    uint
}
