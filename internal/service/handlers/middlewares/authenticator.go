package middlewares

import (
	"net/http"
	"strings"

	"github.com/vldKasatonov/btc-indexer-svc/internal/service/handlers"
	"github.com/vldKasatonov/btc-indexer-svc/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

const authorizationHeader = "Authorization"

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authorizationHeader)
		if authHeader == "" {
			handlers.Log(r).Info("authorization header is missing")
			ape.RenderErr(w, problems.Unauthorized())
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			handlers.Log(r).Info("invalid authorization header")
			ape.RenderErr(w, problems.Unauthorized())
			return
		}

		userId, err := helpers.ValidateJwt(headerParts[1], handlers.Signer(r).JwtSecret)
		if err != nil {
			handlers.Log(r).Info("failed to validate jwt")
			ape.RenderErr(w, problems.Unauthorized())
			return
		}

		next.ServeHTTP(w, r.WithContext(handlers.CtxUserId(userId)(r.Context())))
	})
}
