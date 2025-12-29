package service

import (
	"github.com/go-chi/chi"
	"github.com/vldKasatonov/btc-indexer-svc/internal/data/pg"
	"github.com/vldKasatonov/btc-indexer-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxUsersQ(pg.NewUsersQ(s.config.DB())),
			handlers.CtxSigner(s.config.SignerConfig()),
		),
	)
	r.Route("/integrations/btc-indexer-svc", func(r chi.Router) {
		r.Post("/register", handlers.RegisterUser)
		r.Post("/login", handlers.LoginUser)
	})

	return r
}
