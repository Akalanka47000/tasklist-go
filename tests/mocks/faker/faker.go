// A wrapper around the jaswdr/faker library to extend its functionality
package faker

import (
	"github.com/jaswdr/faker/v2"
	"github.com/samber/lo"
)

type ExtendedFaker struct {
	faker.Faker
}

// Generates a mock Task ID in the format "task_<random_string>"
//
// Note that this is just to demonstrate extending the faker library and does not serve any real purpose.
func (f ExtendedFaker) TaskID() string {
	return "task_" + lo.RandomString(10, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
}

// Global extended faker instance for generating mock data
var Instance = ExtendedFaker{faker.New()}
