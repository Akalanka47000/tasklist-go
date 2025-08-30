package service

import (
	"context"
	"strings"
	. "tasklist/modules/users/api/v1/models"
	repository "tasklist/modules/users/api/v1/repository/contracts"
	"tasklist/modules/users/api/v1/service/contracts"
	"tasklist/utils/hash"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"github.com/samber/lo"
)

// service implements the Service interface
type service struct {
	repo repository.Repository
}

func new(params Params) contracts.Service {
	return &service{
		repo: params.Repository,
	}
}

func (s *service) CreateUser(ctx context.Context, user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(hash.MustString(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *service) GetUsers(ctx context.Context, fqr fq.Result) elemental.PaginateResult[User] {
	return s.repo.GetUsers(ctx, fqr)
}

func (s *service) GetUserByID(ctx context.Context, id string) *User {
	return s.repo.GetUserByID(ctx, id)
}

func (s *service) GetUserByEmail(ctx context.Context, email string) *User {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *service) UpdateUserByID(ctx context.Context, id string, user User) User {
	if user.Password != nil {
		user.Password = lo.ToPtr(hash.MustString(*user.Password))
	}
	if user.Email != nil {
		user.Email = lo.ToPtr(strings.ToLower(*user.Email))
	}
	return s.repo.UpdateUserByID(ctx, id, user)
}

func (s *service) DeleteUserByID(ctx context.Context, id string) User {
	return s.repo.DeleteUserByID(ctx, id)
}
