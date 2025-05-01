package entity

import "time"

type Bid struct {
	Result    bool      `gorm:"index"`
	Bidder    string    `gorm:"index"`
	Course    uint      `gorm:"index"`
	TxHash    string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Amount    uint
}
