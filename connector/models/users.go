package models

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
)

type (
	UpdateUserParams struct {
		Id int64 `json:"id"`
		resources.UserAttributes
	}

	GetUsersRequest requests.GetUsersRequest
)
