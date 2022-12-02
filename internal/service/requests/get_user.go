package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

type GetUserRequest struct {
	UserId int64
}

func NewGetUserRequest(r *http.Request) (GetUserRequest, error) {
	request := GetUserRequest{}
	userId := r.URL.Query().Get("id")
	if userId == "" {
		return GetUserRequest{}, errors.New("failed to get user id")
	}

	var err error
	request.UserId, err = strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return GetUserRequest{}, err
	}

	return request, nil
}
