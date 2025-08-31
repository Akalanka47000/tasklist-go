// Aggregates all mock objects defined within mock files ending with '_mocks.go' and holds the mock values which are used
// across tests. This does not define any mock objects by itself. All definitions should be close to the actual code they mock.
package mocks

import (
	"dario.cat/mergo"
	"tasklist/tests/mocks/faker"

	"github.com/samber/lo"
)

// Utility function to override mock values with provided overrides
// Note that only the first override is used, if multiple are provided.
func mustOverrideOrDefault[T any](base T, overrides ...T) T {
	lo.Must0(mergo.Merge(&base, lo.FirstOrEmpty(overrides), mergo.WithOverride))
	return base
}

// Re-export of the global extended faker instance for generating mock data for easy access in tests.
var Faker = faker.Instance
