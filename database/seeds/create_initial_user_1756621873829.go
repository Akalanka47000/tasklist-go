package seeds

import (
	"context"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "tasklist/modules/users/api/v1/models"
	"tasklist/utils/hash"
)

func Up_1756621873829(ctx context.Context, _ *mongo.Database, _ *mongo.Client) { //nolint:staticcheck // ST1003
	UserModel.Create(User{
		Name:     lo.ToPtr("Akalanka Perera"),
		Email:    lo.ToPtr("akalanka@tasklist.io"),
		Password: lo.ToPtr(hash.MustString("Password@123")),
	}).Exec(ctx)
}

func Down_1756621873829(ctx context.Context, _ *mongo.Database, _ *mongo.Client) { //nolint:staticcheck // ST1003
	UserModel.DeleteOne(primitive.M{"email": "akalanka@tasklist.io"}).Exec(ctx)
}
