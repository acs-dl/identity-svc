package handlers

import (
	"net/http"

	"github.com/acs-dl/identity-svc/internal/data"
	"github.com/acs-dl/identity-svc/internal/service/api/requests"
	"github.com/acs-dl/identity-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUsersRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to unmarshal response")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	users, err := UsersQ(r).Select(data.UserSelector{
		OffsetParams: &request.OffsetPageParams,
		Name:         request.Name,
		Surname:      request.Surname,
		Position:     request.Position,
		Email:        request.Email,
		Search:       request.Search,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to select users")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	totalCount, err := UsersQ(r).Select(data.UserSelector{
		Name:     request.Name,
		Surname:  request.Surname,
		Position: request.Position,
		Email:    request.Email,
		Search:   request.Search,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to select total count")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := newUserListResponse(users, int64(len(totalCount)))
	response.Links = data.GetOffsetLinksForPGParams(r, request.OffsetPageParams)

	ape.Render(w, response)
}

func newUserListResponse(users []data.User, totalCount int64) UserListResponse {
	var response = UserListResponse{
		Meta: Meta{
			TotalCount: totalCount,
		},
		Data: make([]resources.User, len(users)),
	}

	for i, user := range users {
		response.Data[i] = newUserResource(user)
	}

	return response
}

func newUserResource(user data.User) resources.User {
	return resources.User{
		Key: resources.NewKeyInt64(user.Id, resources.USER),
		Attributes: resources.UserAttributes{
			Name:     user.Name,
			Position: user.Position,
			Surname:  user.Surname,
			Email:    user.Email,
			Telegram: user.Telegram,
		},
	}
}

type UserListResponse struct {
	Meta  Meta             `json:"meta"`
	Data  []resources.User `json:"data"`
	Links *resources.Links `json:"links"`
}

type Meta struct {
	TotalCount int64 `json:"total_count"`
}
