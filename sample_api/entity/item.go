package entity

import "time"

type Item struct {
	ID        int       `gorm:"primarykey;autoIncrement"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
