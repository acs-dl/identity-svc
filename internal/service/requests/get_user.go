package requests

import (
	"net/http"
)

type GetUserRequest struct {
	UserId int64
}

func NewGetUserRequest(r *http.Request) (GetUserRequest, error) {
	request := GetUserRequest{}

	userId, err := RetrieveId(r)
	if err != nil {
		return request, err
	}

	request.UserId = userId

	return request, nil
}
