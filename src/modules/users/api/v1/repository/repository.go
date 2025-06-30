package repository

import (
	"tasklist/src/modules/users/api/v1/errors"
	. "tasklist/src/modules/users/api/v1/models"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user User) User {
	return UserModel.Create(user).ExecT()
}

func GetUserByEmail(email string) *User {
	return UserModel.FindOne(primitive.M{"email": email}).ExecPtr()
}

func GetUsers(fqr fq.Result) elemental.PaginateResult[User] {
	fqr.Select["password"] = 0
	return UserModel.QSR(fqr).ExecTP()
}

func GetUserByID(id string, plain ...bool) *User {
	q := UserModel.FindByID(id)
	if len(plain) == 0 || !plain[0] {
		q = q.Select("-password")
	}
	return q.ExecPtr()
}

func UpdateUserByID(id string, user User) User {
	return UserModel.FindByIDAndUpdate(id, user).New().OrFail(errors.UserNotFound).ExecT()
}

func DeleteUserByID(id string) User {
	return UserModel.FindByIDAndDelete(id).OrFail(errors.UserNotFound).ExecT()
}
