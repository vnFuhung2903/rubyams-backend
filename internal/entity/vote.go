package entity

import "time"

type Vote struct {
	Score     uint
	Voter     string    `gorm:"index"`
	Course    uint      `gorm:"index"`
	TxHash    string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
