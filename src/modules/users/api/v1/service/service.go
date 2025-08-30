package service

import (
	"context"
	"strings"
	. "tasklist/src/modules/users/api/v1/models"
	"tasklist/src/modules/users/api/v1/repository"
	"tasklist/src/modules/users/api/v1/service/contracts"
	"tasklist/src/utils/hash"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"github.com/samber/lo"
)

type Service = contracts.Service // Service defines the contract for user service operations

// service implements the Service interface
type service struct {
	repo repository.Repository
}

// New creates a new instance of the user service
func New(params Params) Service {
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
