package handlers

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteUserRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := UsersQ(r).FilterById(request.UserId).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get users")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if user == nil {
		Log(r).Error("user not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	err = UsersQ(r).Delete(request.UserId)
	if err != nil {
		Log(r).WithError(err).Error("failed to delete user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
