package dto

import . "tasklist/modules/users/api/v1/models"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" messages:"Please provide a valid email address"`
	Password string `json:"password" validate:"required" messages:"Please provide a password" mask:"filled"`
}

type LoginResponse = User

type RegisterRequest struct {
	Name     string `json:"name" validate:"required" messages:"Please tell us your name"`
	Email    string `json:"email" validate:"required,email" messages:"Please provide a valid email address"`
	Password string `json:"password" validate:"required,password" messages:"Password should have at least one lowercase letter, one uppercase letter, one number, one special character and should be at least 8 characters long" mask:"filled"`
}

type RegisterResponse = User
