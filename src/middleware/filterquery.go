package middleware

import fqm "github.com/elcengine/elemental/plugins/filterquery/middleware"

var FilterQuery = fqm.NewGoFiber(fqm.Options{
	DefaultLimit: 10,
})
