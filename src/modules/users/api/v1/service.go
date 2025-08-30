package v1

import (
	"context"
	"strings"
	. "tasklist/src/modules/users/api/v1/models"
	"tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/utils/hash"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"github.com/samber/lo"
)

func CreateUser(ctx context.Context, user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(hash.MustString(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return repository.CreateUser(ctx, user)
}

func GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User] {
	return repository.GetUsers(ctx, fqr)
}

func GetUserByID(ctx context.Context, id string) *User {
	return repository.GetUserByID(ctx, id)
}

func UpdateUserByID(ctx context.Context, id string, user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(hash.MustString(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return repository.UpdateUserByID(ctx, id, user)
}

func DeleteUserByID(ctx context.Context, id string) User {
	return repository.DeleteUserByID(ctx, id)
}
