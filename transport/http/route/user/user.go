package user

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/transport/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/clean-architecture/endpoints"
	userDecode "github.com/phungvandat/clean-architecture/transport/http/decode/user"
	"github.com/phungvandat/clean-architecture/transport/http/encode"
)

// UserRoute route
func UserRoute(
	endpoints endpoints.Endpoints,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Find by ID
		r.Get("/{userID}", httptransport.NewServer(
			endpoints.UserEndpoint.FindByID,
			userDecode.FindByIDRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Find
		r.Get("/", httptransport.NewServer(
			endpoints.UserEndpoint.Find,
			userDecode.FindRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		// Create
		r.Post("/", http.NewServer(
			endpoints.UserEndpoint.Create,
			userDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
