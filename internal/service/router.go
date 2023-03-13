package service

import (
	"github.com/go-chi/chi"
	auth "gitlab.com/distributed_lab/acs/auth/middlewares"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	secret := s.config.JwtParams().Secret

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxUsersQ(postgres.NewUsersQ(s.config.DB())),
			handlers.CtxPositions(s.config.Positions()),
		),
	)
	r.Route("/integrations/identity-svc", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.With(auth.Jwt(secret, "identity", []string{"read", "write"}...)).
				Get("/", handlers.GetUsers)
			r.With(auth.Jwt(secret, "identity", []string{"write"}...)).
				Post("/", handlers.CreateUser)

			r.With(auth.Jwt(secret, "identity", []string{"read", "write"}...)).
				Get("/{id}", handlers.GetUser)
			r.With(auth.Jwt(secret, "identity", []string{"write"}...)).
				Delete("/{id}", handlers.DeleteUser)
			r.With(auth.Jwt(secret, "identity", []string{"write"}...)).
				Patch("/{id}", handlers.UpdateUser)

			r.With(auth.Jwt(secret, "identity", []string{"read", "write"}...)).
				Get("/positions", handlers.GetPositions)
		})
	})

	return r
}
