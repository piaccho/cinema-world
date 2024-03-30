package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=64"`
	Firstname string `json:"firstname" binding:"required,min=2"`
	Lastname  string `json:"lastname" binding:"required,min=2"`
}
