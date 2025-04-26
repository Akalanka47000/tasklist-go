package v1

import (
	"tasklist/src/modules/users/api/v1/models"
	"tasklist/src/utils"
)

var repository = utils.NewRepository[models.User]("users")

func Repository() utils.Repository[models.User] {
	return repository
}
