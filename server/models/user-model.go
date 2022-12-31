package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique" `
	Username string `json:"username" `
	Password string `json:"password"`
	Image    string `json:"image"`
}
