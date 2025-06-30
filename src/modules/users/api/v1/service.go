package v1

import (
	"strings"
	. "tasklist/src/modules/users/api/v1/models"
	"tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/utils"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"github.com/samber/lo"
)

func CreateUser(user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(utils.HashStr(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return repository.CreateUser(user)
}

func GetUsers(fqr fq.Result) elemental.PaginateResult[User] {
	return repository.GetUsers(fqr)
}

func GetUserByID(id string) *User {
	return repository.GetUserByID(id)
}

func UpdateUserByID(id string, user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(utils.HashStr(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return repository.UpdateUserByID(id, user)
}

func DeleteUserByID(id string) User {
	return repository.DeleteUserByID(id)
}
