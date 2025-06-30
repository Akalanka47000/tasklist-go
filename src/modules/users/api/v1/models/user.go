package models

import (
	"reflect"
	"time"

	elemental "github.com/elcengine/elemental/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      *string            `json:"name,omitempty" bson:"name,omitempty"`
	Email     *string            `json:"email,omitempty" bson:"email,omitempty"`
	Password  *string            `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

var UserModel = elemental.NewModel[User]("User", elemental.NewSchema(map[string]elemental.Field{
	"Name": {
		Type: reflect.String,
	},
	"Email": {
		Type: reflect.String,
		Index: options.Index().SetUnique(true).
			SetPartialFilterExpression(primitive.M{"email": primitive.M{"$exists": true}}),
	},
	"Password": {
		Type: reflect.String,
	},
}, elemental.SchemaOptions{
	Auditing: true,
}))
