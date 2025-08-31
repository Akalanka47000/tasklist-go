// Aggregates all mock objects defined within mock files ending with '_mocks.go' and holds the mock values which are used
// across tests. This does not define any mock objects by itself. All definitions should be close to the actual code they mock.
package mocks

import "tasklist/tests/mocks/faker"

// Reexport of the global extended faker instance for generating mock data for easy access in tests.
var Faker = faker.Instance
