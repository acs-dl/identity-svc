package models

import (
	"github.com/acs-dl/identity-svc/internal/service/api/requests"
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
