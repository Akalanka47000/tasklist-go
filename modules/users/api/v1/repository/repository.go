package repository

import (
	"context"
	"tasklist/modules/users/api/v1/errors"
	. "tasklist/modules/users/api/v1/models"
	"tasklist/modules/users/api/v1/repository/contracts"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository struct{} // repository implements UserRepository

func new() contracts.Repository {
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
	return UserModel.FindByIDAndUpdate(id, user).New().OrFail(httperrs.UserNotFound).ExecT(ctx)
}

func (r *repository) DeleteUserByID(ctx context.Context, id string) User {
	return UserModel.FindByIDAndDelete(id).OrFail(httperrs.UserNotFound).ExecT(ctx)
}
