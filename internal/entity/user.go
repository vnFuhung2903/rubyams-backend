package entity

type User struct {
	ID      uint    `gorm:"primaryKey"`
	Email   string  `gorm:"unique"`
	Name    string  `gorm:"not null"`
	Hash    string  `gorm:"not null"`
	Student Student `gorm:"foreignKey:UserID"`
}
