package handlers

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/data"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSearchUsersRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to unmarshal response")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	users, err := UsersQ(r).SearchBy(*request.Search).Select(data.UserSelector{OffsetParams: &request.OffsetPageParams})
	if err != nil {
		Log(r).WithError(err).Error("failed to select users")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	totalCount, err := UsersQ(r).SearchBy(*request.Search).Select(data.UserSelector{})
	if err != nil {
		Log(r).WithError(err).Error("failed to select total count")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := newUserListResponse(users, int64(len(totalCount)))
	response.Links = data.GetOffsetLinksForPGParams(r, request.OffsetPageParams)

	ape.Render(w, response)
}
