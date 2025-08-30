package repository

import (
	"context"
	"tasklist/src/modules/users/api/v1/errors"
	. "tasklist/src/modules/users/api/v1/models"
	"tasklist/src/modules/users/api/v1/repository/contracts"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository = contracts.Repository // Repository defines the contract for user repository operations

type repository struct{} // repository implements UserRepository

// New creates a new instance of UserRepository
func New() Repository {
	return &repository{}
}

func (r *repository) CreateUser(ctx context.Context, user User) User {
	return UserModel.Create(user).ExecT(ctx)
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) *User {
	return UserModel.FindOne(primitive.M{"email": email}).ExecPtr(ctx)
}

func (r *repository) GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User] {
	fqr.Select["password"] = 0
	return UserModel.QSR(fqr).ExecTP(ctx)
}

func (r *repository) GetUserByID(ctx context.Context, id string, plain ...bool) *User {
	q := UserModel.FindByID(id)
	if len(plain) == 0 || !plain[0] {
		q = q.Select("-password")
	}
	return q.ExecPtr(ctx)
}

func (r *repository) UpdateUserByID(ctx context.Context, id string, user User) User {
	return UserModel.FindByIDAndUpdate(id, user).New().OrFail(errors.UserNotFound).ExecT(ctx)
}

func (r *repository) DeleteUserByID(ctx context.Context, id string) User {
	return UserModel.FindByIDAndDelete(id).OrFail(errors.UserNotFound).ExecT(ctx)
}
