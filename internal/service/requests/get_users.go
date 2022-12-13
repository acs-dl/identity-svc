package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetUsersRequest struct {
	pgdb.OffsetPageParams

	Name     *string `filter:"name"`
	Surname  *string `filter:"surname"`
	Position *string `filter:"position"`
}

func NewGetUsersRequest(r *http.Request) (GetUsersRequest, error) {
	request := GetUsersRequest{}

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return request, errors.Wrap(err, "failed to decode request")
	}

	return request, nil
}
