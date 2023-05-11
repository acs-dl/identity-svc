package connector

import (
	"github.com/acs-dl/identity-svc/connector/models"
	"github.com/acs-dl/identity-svc/internal/service/api/handlers"
	"github.com/acs-dl/identity-svc/resources"
)

type ConnectorI interface {
	CreateUser(user resources.User) (resources.UserResponse, error)
	GetUser(userId int64) (resources.UserResponse, error)
	DeleteUser(userId int64) error
	UpdateUser(request models.UpdateUserParams) error
	GetUsers(params models.GetUsersRequest) (handlers.UserListResponse, error)

	GetPositions() (resources.PositionsResponse, error)
}
