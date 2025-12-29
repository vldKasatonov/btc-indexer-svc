package helpers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vldKasatonov/btc-indexer-svc/resources"
)

func ValidateCredentials(credentials resources.UserCredentials) validation.Errors {
	return validation.Errors{
		"/data/attributes/username": validation.Validate(
			&credentials.Attributes.Username,
			validation.Required,
			validation.Length(4, 128),
		),
		"/data/attributes/password": validation.Validate(
			&credentials.Attributes.Password,
			validation.Required,
			validation.Length(8, 128),
		),
	}
}
