package mocks

import (
	. "tasklist/modules/users/api/v1/models"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(overrides ...User) User {
	return mustOverrideOrDefault(User{
		ID:       primitive.NewObjectID(),
		Name:     lo.ToPtr(Faker.Person().Name()),
		Email:    lo.ToPtr(Faker.Internet().Email()),
		Password: lo.ToPtr(Faker.Internet().Password()),
	}, overrides...)
}
