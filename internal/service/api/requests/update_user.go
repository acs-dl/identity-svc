package requests

import (
	"encoding/json"
	"net/http"

	"github.com/acs-dl/identity-svc/resources"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type UpdateUserRequest struct {
	Data   resources.User `json:"data"`
	UserID int64
}

func NewUpdateUserRequest(r *http.Request) (UpdateUserRequest, error) {
	request := UpdateUserRequest{
		UserID: cast.ToInt64(chi.URLParam(r, "id")),
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, validate(request.Data)

}
