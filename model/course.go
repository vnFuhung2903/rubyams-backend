package model

type Course struct {
	Id          uint       `gorm:"primaryKey"`
	CourseName  string     `gorm:"not null;index"`
	TeacherName string     `gorm:"not null;index"`
	Semester    uint       `gorm:"index"`
	Students    []*Student `gorm:"many2many:students_classes"`
	InProgress  bool       `gorm:"index"`
	Date        string     `gorm:"index"`
	StartTime   string
	EndTime     string
	Capacity    uint
}
