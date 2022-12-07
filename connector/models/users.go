package models

import "gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"

type (
	UpdateUserParams struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Position string `json:"position"`
	}

	GetUsersRequest requests.GetUsersRequest
)
