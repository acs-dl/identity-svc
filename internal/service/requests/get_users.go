package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type GetUsersRequest struct {
	pgdb.OffsetPageParams
}

func NewGetUsersRequest(r *http.Request) (GetUsersRequest, error) {
	request := GetUsersRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, nil
}
