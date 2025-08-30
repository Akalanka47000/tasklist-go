package contracts

import (
	"context"
	. "tasklist/src/modules/users/api/v1/models"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
)

// Service defines the contract for user service operations
type Service interface {
	// CreateUser creates a new user with password hashing and email normalization
	CreateUser(ctx context.Context, user User) User
	// GetUsers retrieves a paginated list of users based on filter query
	GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User]
	// GetUserByID retrieves a user by their ID
	GetUserByID(ctx context.Context, id string) *User
	// GetUserByEmail retrieves a user by their email address
	GetUserByEmail(ctx context.Context, email string) *User
	// UpdateUserByID updates a user by their ID with password hashing and email normalization
	UpdateUserByID(ctx context.Context, id string, user User) User
	// DeleteUserByID deletes a user by their ID
	DeleteUserByID(ctx context.Context, id string) User
}
