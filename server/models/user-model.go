package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3"`
	Image    string `json:"image" validate:"required"`
}
