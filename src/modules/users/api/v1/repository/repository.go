package repository

import (
	"context"
	"tasklist/src/modules/users/api/v1/errors"
	. "tasklist/src/modules/users/api/v1/models"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(ctx context.Context, user User) User {
	return UserModel.Create(user).ExecT(ctx)
}

func GetUserByEmail(ctx context.Context, email string) *User {
	return UserModel.FindOne(primitive.M{"email": email}).ExecPtr()
}

func GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User] {
	fqr.Select["password"] = 0
	return UserModel.QSR(fqr).ExecTP(ctx)
}

func GetUserByID(ctx context.Context, id string, plain ...bool) *User {
	q := UserModel.FindByID(id)
	if len(plain) == 0 || !plain[0] {
		q = q.Select("-password")
	}
	return q.ExecPtr(ctx)
}

func UpdateUserByID(ctx context.Context, id string, user User) User {
	return UserModel.FindByIDAndUpdate(id, user).New().OrFail(errors.UserNotFound).ExecT(ctx)
}

func DeleteUserByID(ctx context.Context, id string) User {
	return UserModel.FindByIDAndDelete(id).OrFail(errors.UserNotFound).ExecT(ctx)
}
