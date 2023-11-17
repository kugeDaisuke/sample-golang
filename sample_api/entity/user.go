package entity

type User struct {
	ID    int    `gorm:"primarykey;autoIncrement"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Pass  string `gorm:"not null"`
}
