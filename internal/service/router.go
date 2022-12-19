package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

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
			r.Get("/", handlers.GetUsers)
			r.Get("/{id}", handlers.GetUser)
			r.Delete("/{id}", handlers.DeleteUser)
			r.Post("/", handlers.CreateUser)
			r.Patch("/{id}", handlers.UpdateUser)

			r.Get("/positions", handlers.GetPositions)
		})
	})

	return r
}
