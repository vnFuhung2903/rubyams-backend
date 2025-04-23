package model

import "time"

type Bid struct {
	Result    bool      `gorm:"index"`
	Voter     string    `gorm:"index"`
	Semester  uint      `gorm:"index"`
	Class     string    `gorm:"index"`
	TxHash    string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Amount    uint
}
