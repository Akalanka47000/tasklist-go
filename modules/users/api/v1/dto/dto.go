package dto

import (
	. "tasklist/modules/users/api/v1/models"

	elemental "github.com/elcengine/elemental/core"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" messages:"Please provide a name for the user"`
	Email    string `json:"email" validate:"required,email" messages:"Please provide a valid email address"`
	Password string `json:"password" validate:"required,password" messages:"Password should have at least one lowercase letter, one uppercase letter, one number, one special character and should be at least 8 characters long" mask:"filled"`
}

type CreateUserResponse = User

type GetUsersReponse = elemental.PaginateResult[User]

type GetUserRequest struct {
	ID string `json:"id" validate:"required,objectid" messages:"Please provide a valid user ID"`
}

type GetUserResponse = User

type UpdateUserRequest struct {
	GetUserRequest
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"omitempty,email" messages:"Please provide a valid email address"`
	Password *string `json:"password" validate:"omitempty,password" messages:"Password should have at least one lowercase letter, one uppercase letter, one number, one special character and should be at least 8 characters long" mask:"filled"`
}

type DeleteUserRequest = GetUserRequest
