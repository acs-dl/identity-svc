package models

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
)

type (
	UpdateUserParams struct {
		Id       int64  `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Surname  string `json:"surname"`
	}

	GetUsersRequest requests.GetUsersRequest
)
