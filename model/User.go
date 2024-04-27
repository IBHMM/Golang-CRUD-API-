package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"size:50;not null" json:"username"`
	Email    string `gorm:"size:100;not null" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}
