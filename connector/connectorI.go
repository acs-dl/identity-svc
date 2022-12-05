package connector

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type ConnectorI interface {
	CreateUser(user resources.User) (resources.UserResponse, error)
	GetUser(userId int64) (resources.UserResponse, error)
	DeleteUser(userId int64) error
	UpdateUser(request requests.UpdateUserRequest) (resources.UserResponse, error)
	GetUsers(params pgdb.OffsetPageParams) (resources.UserListResponse, error)
}
