package service

import (
	"context"
	"strings"
	"testing"

	"tasklist/modules/users/api/v1/models"
	repository "tasklist/modules/users/api/v1/repository/contracts"
	"tasklist/tests/mocks"

	elemental "github.com/elcengine/elemental/core"
	fq "github.com/elcengine/elemental/plugins/filterquery"
	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	Convey("V1", t, func() {
		mockRepo := repository.NewMockRepository(t)
		svc := new(Params{Repository: mockRepo})
		ctx := t.Context()

		mockUser := mocks.GetUser()

		Convey("CreateUser", func() {
			mockRepo.EXPECT().CreateUser(ctx, mock.Anything).RunAndReturn(func(ctx context.Context, user models.User) models.User {
				return user
			})
			Convey("Given valid user data, should hash password, lowercase email, and return created user", func() {
				mockUser.Email = lo.ToPtr(strings.ToUpper(*mockUser.Email))
				result := svc.CreateUser(ctx, mockUser)
				So(*result.Email, ShouldEqual, strings.ToLower(*mockUser.Email))
				So(*result.Password, ShouldNotEqual, *mockUser.Password)
				So(*result.Password, ShouldStartWith, "$2a$") // bcrypt hash prefix
			})
			Convey("Given user data with nil password, should not attempt to hash and return created user", func() {
				mockUser.Password = nil
				result := svc.CreateUser(ctx, mockUser)
				So(result.Password, ShouldBeNil)
			})
		})

		Convey("GetUsers", func() {
			Convey("should return an empty result when no users exist", func() {
				fqr := fq.Result{}
				paginatedResult := elemental.PaginateResult[models.User]{}
				mockRepo.EXPECT().GetUsers(ctx, fqr).Return(paginatedResult)
				result := svc.GetUsers(ctx, fqr)
				So(result, ShouldEqual, paginatedResult)
			})
			Convey("should return a paginated list of users when users exist", func() {
				fqr := fq.Result{}
				users := []models.User{
					mocks.GetUser(),
					mocks.GetUser(),
				}
				paginatedResult := elemental.PaginateResult[models.User]{
					Docs:       users,
					TotalDocs:  int64(len(users)),
					Limit:      10,
					Page:       1,
					TotalPages: 1,
				}
				mockRepo.EXPECT().GetUsers(ctx, fqr).Return(paginatedResult)
				result := svc.GetUsers(ctx, fqr)
				So(result, ShouldResemble, paginatedResult)
			})
		})

		Convey("GetUserByID", func() {
			Convey("should return the user when found", func() {
				mockRepo.EXPECT().GetUserByID(ctx, mockUser.ID.Hex()).Return(&mockUser)
				result := svc.GetUserByID(ctx, mockUser.ID.Hex())
				So(result, ShouldResemble, &mockUser)
			})
			Convey("should return nil when user is not found", func() {
				mockRepo.EXPECT().GetUserByID(ctx, mockUser.ID.Hex()).Return(nil)
				result := svc.GetUserByID(ctx, mockUser.ID.Hex())
				So(result, ShouldBeNil)
			})
		})

		Convey("GetUserByEmail", func() {
			Convey("should return the user when found", func() {
				mockRepo.EXPECT().GetUserByEmail(ctx, *mockUser.Email).Return(&mockUser)
				result := svc.GetUserByEmail(ctx, *mockUser.Email)
				So(result, ShouldResemble, &mockUser)
			})
			Convey("should return nil when user is not found", func() {
				mockRepo.EXPECT().GetUserByEmail(ctx, *mockUser.Email).Return(nil)
				result := svc.GetUserByEmail(ctx, *mockUser.Email)
				So(result, ShouldBeNil)
			})
		})

		Convey("UpdateUserByID", func() {
			mockRepo.EXPECT().UpdateUserByID(ctx, mockUser.ID.Hex(), mock.Anything).RunAndReturn(func(ctx context.Context, id string, user models.User) models.User {
				return user
			})
			Convey("should hash password, lowercase email, and return updated user", func() {
				result := svc.UpdateUserByID(ctx, mockUser.ID.Hex(), mockUser)
				So(result.Email, ShouldResemble, mockUser.Email)
				So(result.Password, ShouldNotResemble, mockUser.Password)
				So(*result.Password, ShouldStartWith, "$2a$") // bcrypt hash prefix
			})
			Convey("should not hash password if nil and should lowercase email", func() {
				mockUser.Password = nil
				result := svc.UpdateUserByID(ctx, mockUser.ID.Hex(), mockUser)
				So(result.Email, ShouldResemble, mockUser.Email)
				So(result.Password, ShouldBeNil)
			})
		})

		Convey("DeleteUserByID", func() {
			Convey("should return the deleted user", func() {
				mockRepo.EXPECT().DeleteUserByID(ctx, mockUser.ID.Hex()).Return(mockUser)
				result := svc.DeleteUserByID(ctx, mockUser.ID.Hex())
				So(result, ShouldResemble, mockUser)
			})
			Convey("should return zero value user if not found", func() {
				mockRepo.EXPECT().DeleteUserByID(ctx, mockUser.ID.Hex()).Return(models.User{})
				result := svc.DeleteUserByID(ctx, mockUser.ID.Hex())
				So(result, ShouldBeZeroValue)
			})
		})
	})
}
