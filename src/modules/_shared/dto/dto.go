// Shared DTOs between modules
package sdto

type PaginatedQuery struct {
	Page  int `query:"page" validate:"min=1" messages:"You must provide a valid page number to use this endpoint (minimum is 1)"`
	Limit int `query:"limit" validate:"omitempty,min=1,max=100" messages:"Limit should be between 1 and 100"`
}
