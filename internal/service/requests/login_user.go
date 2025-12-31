package requests

import (
	"encoding/json"
	"net/http"

	"github.com/vldKasatonov/btc-indexer-svc/internal/service/helpers"
	"github.com/vldKasatonov/btc-indexer-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LoginUserRequest struct {
	Data resources.UserCredentials `json:"data"`
}

func NewLoginUserRequest(r *http.Request) (LoginUserRequest, error) {
	var request LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, request.validate()
}

func (r *LoginUserRequest) validate() error {
	return helpers.ValidateCredentials(r.Data).Filter()
}
