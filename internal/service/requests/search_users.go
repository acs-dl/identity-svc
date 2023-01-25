package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type SearchUsersRequest struct {
	pgdb.OffsetPageParams

	Search *string `filter:"search"`
}

func NewSearchUsersRequest(r *http.Request) (SearchUsersRequest, error) {
	request := SearchUsersRequest{}

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
