package repository

import (
	"testing"

	"tasklist/tests/mocks"

	// fq "github.com/elcengine/elemental/plugins/filterquery"
	elemental "github.com/elcengine/elemental/core"
	. "github.com/smartystreets/goconvey/convey"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	repo := new()
	mockUser := mocks.GetUser()

	Convey("CreateUser", t, func() {
		mt.Run("CreateUser", func(mt *mtest.T) {
			firstBatch := bson.D{
				{Key: "_id", Value: bson.D{{Key: "ts", Value: 1}}},
				{Key: "operationType", Value: "insert"},
				{Key: "fullDocument", Value: bson.D{{Key: "name", Value: "Alice"}}},
			}

			mt.AddMockResponses(
				mtest.CreateSuccessResponse(bson.E{
					Key: "insertedIds", Value: bson.A{"mocked-id"},
				}),
				mtest.CreateCursorResponse(1, "db.coll", mtest.FirstBatch, firstBatch),
			)
			elemental.Connect(mt.Client)
			result := repo.CreateUser(mt.Context(), mockUser)
			So(result, ShouldResemble, mockUser)
		})
	})

	// Convey("GetUserByEmail", t, func() {
	// 	userDoc := bson.D{
	// 		bson.E{Key: "_id", Value: mockUser.ID},
	// 		bson.E{Key: "name", Value: mockUser.Name},
	// 		bson.E{Key: "email", Value: mockUser.Email},
	// 		bson.E{Key: "password", Value: mockUser.Password},
	// 	}
	// 	mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch, userDoc))
	// 	result := repo.GetUserByEmail(mt.Context(), *mockUser.Email)
	// 	So(result, ShouldResemble, &mockUser)
	// })

	// Convey("GetUsers", t, func() {
	// 	userDoc := bson.D{
	// 		bson.E{Key: "_id", Value: mockUser.ID},
	// 		bson.E{Key: "name", Value: mockUser.Name},
	// 		bson.E{Key: "email", Value: mockUser.Email},
	// 		bson.E{Key: "password", Value: mockUser.Password},
	// 	}
	// 	mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch, userDoc))
	// 	fqr := fq.Result{}
	// 	result := repo.GetUsers(mt.Context(), fqr)
	// 	So(len(result.Docs), ShouldEqual, 1)
	// 	So(result.Docs[0], ShouldResemble, mockUser)
	// })

	// Convey("GetUserByID", t, func() {
	// 	userDoc := bson.D{
	// 		bson.E{Key: "_id", Value: mockUser.ID},
	// 		bson.E{Key: "name", Value: mockUser.Name},
	// 		bson.E{Key: "email", Value: mockUser.Email},
	// 		bson.E{Key: "password", Value: mockUser.Password},
	// 	}
	// 	mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch, userDoc))
	// 	result := repo.GetUserByID(mt.Context(), mockUser.ID.Hex())
	// 	So(result, ShouldResemble, &mockUser)
	// })

	// Convey("UpdateUserByID", t, func() {
	// 	mt.AddMockResponses(mtest.CreateSuccessResponse())
	// 	result := repo.UpdateUserByID(mt.Context(), mockUser.ID.Hex(), mockUser)
	// 	So(result, ShouldResemble, mockUser)
	// })

	// Convey("DeleteUserByID", t, func() {
	// 	mt.AddMockResponses(mtest.CreateSuccessResponse())
	// 	result := repo.DeleteUserByID(mt.Context(), mockUser.ID.Hex())
	// 	So(result, ShouldResemble, mockUser)
	// })
}
