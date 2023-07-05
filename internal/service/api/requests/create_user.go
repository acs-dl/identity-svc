package requests

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/acs-dl/identity-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var emailRegex = regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")

func NewCreateUserRequest(r *http.Request) (resources.User, error) {
	request := struct {
		Data resources.User
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to decode request")
	}

	return request.Data, validate(request.Data)

}

func validate(r resources.User) error {
	return validation.Errors{
		"/data/attributes/name":     validation.Validate(r.Attributes.Name, validation.Required),
		"/data/attributes/surname":  validation.Validate(r.Attributes.Surname, validation.Required),
		"/data/attributes/position": validation.Validate(r.Attributes.Position, validation.Required),
	}.Filter()
}
