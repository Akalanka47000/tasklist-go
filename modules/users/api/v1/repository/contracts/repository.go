package contracts

import (
	"context"
	. "tasklist/modules/users/api/v1/models"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
)

// Repository defines the contract for user repository operations
type Repository interface {
	// CreateUser creates a new user in the database
	CreateUser(ctx context.Context, user User) User
	// GetUserByEmail retrieves a user by their email address
	GetUserByEmail(ctx context.Context, email string) *User
	// GetUsers retrieves a paginated list of users based on filter query
	GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User]
	// GetUserByID retrieves a user by their ID
	GetUserByID(ctx context.Context, id string, plain ...bool) *User
	// UpdateUserByID updates a user by their ID
	UpdateUserByID(ctx context.Context, id string, user User) User
	// DeleteUserByID deletes a user by their ID
	DeleteUserByID(ctx context.Context, id string) User
}
