package handlers

import (
	"net/http"

	"github.com/acs-dl/identity-svc/internal/data"
	"github.com/acs-dl/identity-svc/internal/service/api/requests"
	"github.com/acs-dl/identity-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateUserRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	id, err := UsersQ(r).Create(data.User{
		Name:     request.Attributes.Name,
		Surname:  request.Attributes.Surname,
		Position: request.Attributes.Position,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.UserResponse{
		Data: resources.User{
			Key: resources.NewKeyInt64(id, resources.USER),
			Attributes: resources.UserAttributes{
				Name:     request.Attributes.Name,
				Position: request.Attributes.Position,
				Surname:  request.Attributes.Surname,
			},
		},
	})
}
