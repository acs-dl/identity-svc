package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func NewCreateUserRequest(r *http.Request) (resources.User, error) {
	request := struct {
		Data resources.User
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to decode request")
	}

	return request.Data, nil

}
