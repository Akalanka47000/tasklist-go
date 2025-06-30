// Shared DTOs between modules
package sdto

type PaginatedQuery struct {
	Page  int `query:"page"`
	Limit int `query:"limit" validate:"min=0,max=100" messages:"Limit should be between 0 and 100"`
}
