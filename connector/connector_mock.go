package connector

import (
	"gitlab.com/distributed_lab/acs/identity-svc/connector/models"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
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

func (c *ConnectorMockIdentity) UpdateUser(request models.UpdateUserParams) error {
	return nil
}

func (c *ConnectorMockIdentity) GetUsers(params models.GetUsersRequest) (handlers.UserListResponse, error) {
	return handlers.UserListResponse{}, nil
}
