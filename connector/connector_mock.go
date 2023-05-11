package connector

import (
	"github.com/acs-dl/identity-svc/connector/models"
	"github.com/acs-dl/identity-svc/internal/service/api/handlers"
	"github.com/acs-dl/identity-svc/resources"
)

type ConnectorMockIdentity struct {
	ServiceUrl string
}

func NewConnectorMockIdentity(serviceUrl string) ConnectorI {
	return &ConnectorMockIdentity{
		ServiceUrl: serviceUrl,
	}
}

func (c *ConnectorMockIdentity) GetPositions() (resources.PositionsResponse, error) {
	return resources.PositionsResponse{}, nil
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
