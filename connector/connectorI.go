package connector

import (
	"gitlab.com/distributed_lab/acs/identity-svc/connector/models"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
)

type ConnectorI interface {
	CreateUser(user resources.User) (resources.UserResponse, error)
	GetUser(userId int64) (resources.UserResponse, error)
	DeleteUser(userId int64) error
	UpdateUser(request models.UpdateUserParams) error
	GetUsers(params models.GetUsersRequest) (handlers.UserListResponse, error)

	GetPositions() (resources.PositionsResponse, error)
}
