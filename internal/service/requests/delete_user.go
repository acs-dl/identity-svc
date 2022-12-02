package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

type DeleteUserRequest struct {
	UserId int64
}

func NewDeleteUserRequest(r *http.Request) (DeleteUserRequest, error) {
	request := DeleteUserRequest{}
	userId := r.URL.Query().Get("id")
	if userId == "" {
		return DeleteUserRequest{}, errors.New("failed to get user id")
	}

	var err error
	request.UserId, err = strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return DeleteUserRequest{}, err
	}

	return request, nil
}
