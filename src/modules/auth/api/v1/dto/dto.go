package dto

import . "tasklist/src/modules/users/api/v1/models"

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginResponse = User

type RegisterRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type RegisterResponse = User
