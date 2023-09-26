package handlers

import (
	"net/http"

	"github.com/acs-dl/identity-svc/internal/data"
	"github.com/acs-dl/identity-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateUserRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := UsersQ(r).GetById(request.UserID)
	if err != nil {
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if user == nil {
		Log(r).Debug("user not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	err = UsersQ(r).Update(data.User{
		Id:       user.Id,
		Name:     request.Data.Attributes.Name,
		Surname:  request.Data.Attributes.Surname,
		Position: request.Data.Attributes.Position,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to update user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
