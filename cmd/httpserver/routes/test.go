package routes

import (
	"net/http"
	"test/internal/handlers/testhttphdl"
)

func NewTestRoutes(th *testhttphdl.HTTPHandler) []*Route {
	return []*Route{
		{
			Path:    "/getall",
			Method:  http.MethodGet,
			Handler: th.GetAllTests, // ??
		},
		{
			Path:    "/get/{id}",
			Method:  http.MethodGet,
			Handler: th.GetTest, // ??
		},
	}
}
