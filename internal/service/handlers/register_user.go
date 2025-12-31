package handlers

import (
	"net/http"

	"github.com/vldKasatonov/btc-indexer-svc/internal/data"
	"github.com/vldKasatonov/btc-indexer-svc/internal/service/helpers"
	"github.com/vldKasatonov/btc-indexer-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRegisterUserRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("failed to decode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	hashedPassword, err := helpers.EncryptPassword(request.Data.Attributes.Password)
	if err != nil {
		Log(r).WithError(err).Info("failed to encrypt password")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	newUser := data.User{
		Username:     request.Data.Attributes.Username,
		PasswordHash: hashedPassword,
	}

	_, err = UsersQ(r).Insert(newUser)
	if err != nil {
		Log(r).WithError(err).Info("failed to create user")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
