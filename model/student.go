package model

type Student struct {
	Id           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Identity     string    `json:"-"`
	Wallet       string    `gorm:"unique"`
	NFTAddress   string    `gorm:"unique"`
	AcademicYear uint      `gorm:"not null; index"`
	Courses      []*Course `gorm:"many2many:students_classes"`
	EnrolledAt   string    `gorm:"not null"`
	GraduatedAt  string
}
