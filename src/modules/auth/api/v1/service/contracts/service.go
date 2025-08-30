package contracts

import (
	"context"
	"tasklist/src/modules/auth/api/v1/dto"
	. "tasklist/src/modules/users/api/v1/models"
)

type Service interface {
	// Login authenticates a user by email and password, returning the user and tokens.
	Login(ctx context.Context, email, password string) (User, string, string)
	// RegisterUser creates a new user and returns the user and tokens.
	RegisterUser(ctx context.Context, payload dto.RegisterRequest) (User, string, string)
}
