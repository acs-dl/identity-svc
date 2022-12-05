package connector

import (
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type ConnectorMockIdentity struct {
	ServiceUrl string
}

func NewConnectorMockIdentity(serviceUrl string) ConnectorI {
	return &ConnectorMockIdentity{
		ServiceUrl: serviceUrl,
	}
}

func (c *ConnectorMockIdentity) CreateUser(user resources.User) (resources.UserResponse, error) {
	return resources.UserResponse{}, nil
}

func (c *ConnectorMockIdentity) GetUser(userId int64) (resources.UserResponse, error) {
	return resources.UserResponse{}, nil
}

func (c *ConnectorMockIdentity) DeleteUser(userId int64) error {
	return nil
}

func (c *ConnectorMockIdentity) UpdateUser(request requests.UpdateUserRequest) (resources.UserResponse, error) {
	return resources.UserResponse{}, nil
}

func (c *ConnectorMockIdentity) GetUsers(params pgdb.OffsetPageParams) (resources.UserListResponse, error) {
	return resources.UserListResponse{}, nil
}
