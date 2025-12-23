package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vldKasatonov/btc-indexer-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RegisterUserRequest struct {
	Data resources.RegisterUser `json:"data"`
}

func NewRegisterUserRequest(r *http.Request) (RegisterUserRequest, error) {
	var request RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, request.validate()
}

func (r *RegisterUserRequest) validate() error {
	return validation.Errors{
		"/data/attributes/username": validation.Validate(
			&r.Data.Attributes.Username,
			validation.Required,
			validation.Length(4, 128),
		),
		"/data/attributes/password": validation.Validate(
			&r.Data.Attributes.Password,
			validation.Required,
			validation.Length(8, 128),
		),
	}.Filter()
}
