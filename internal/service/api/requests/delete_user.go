package requests

import (
	"net/http"
)

type DeleteUserRequest struct {
	UserId int64
}

func NewDeleteUserRequest(r *http.Request) (DeleteUserRequest, error) {
	request := DeleteUserRequest{}

	userId, err := RetrieveId(r)
	if err != nil {
		return request, err
	}

	request.UserId = userId

	return request, nil
}
