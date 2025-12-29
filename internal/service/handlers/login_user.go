package handlers

import (
	"net/http"

	"github.com/vldKasatonov/btc-indexer-svc/internal/service/helpers"
	"github.com/vldKasatonov/btc-indexer-svc/internal/service/requests"
	"github.com/vldKasatonov/btc-indexer-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewLoginUserRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("failed to decode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := UsersQ(r).FilterByUsername(request.Data.Attributes.Username).Get()
	if err != nil {
		Log(r).WithError(err).Info("failed to get user")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}
	err = helpers.VerifyPassword(user.PasswordHash, request.Data.Attributes.Password)
	if err != nil {
		Log(r).WithError(err).Info("invalid password")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	jwt, err := helpers.GenerateJwt(user.ID, Signer(r).JwtSecret)
	if err != nil {
		Log(r).WithError(err).Info("failed to generate jwt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.JwtResponse{
		Data: resources.Jwt{
			Key: resources.NewKeyInt64(0, resources.JWT),
			Attributes: resources.JwtAttributes{
				Token: jwt,
			},
		},
	})
}
