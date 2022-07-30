package models

type User struct{
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Number string `json:"number"`
	Email string `json:"email"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required"`
	Email string `json:"email" binding:"required"`
}
