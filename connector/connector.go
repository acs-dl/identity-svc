package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/acs/identity-svc/connector/models"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/api/handlers"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type Connector struct {
	ServiceUrl string
	Client     *http.Client
}

func NewConnector(serviceUrl string) ConnectorI {
	return &Connector{
		ServiceUrl: serviceUrl,
		Client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Connector) doRequest(method, url string, body interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		response.Body.Close()
		return response, errors.New("Bad status")
	}

	return response, nil
}

func (c *Connector) DoRequest(method, url string, body interface{}) error {
	_, err := c.doRequest(method, url, body)
	return err
}

func (c *Connector) DoRequestWithDecode(method, url string, body, decodeModel interface{}) error {
	response, err := c.doRequest(method, url, body)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(response.Body).Decode(decodeModel); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}

	return nil
}

func (c *Connector) CreateUser(user resources.User) (resources.UserResponse, error) {
	body := struct {
		Data resources.User
	}{Data: user}
	model := resources.UserResponse{}

	err := c.DoRequestWithDecode("POST", c.ServiceUrl+"/users", body, &model)
	return model, err
}

func (c *Connector) GetUser(userId int64) (resources.UserResponse, error) {
	model := resources.UserResponse{}

	err := c.DoRequestWithDecode("GET", fmt.Sprintf("%s/users/%d", c.ServiceUrl, userId), nil, &model)
	return model, err
}

func (c *Connector) DeleteUser(userId int64) error {
	err := c.DoRequest("DELETE", fmt.Sprintf("%s/users/%d", c.ServiceUrl, userId), nil)
	return err
}

func (c *Connector) UpdateUser(request models.UpdateUserParams) error {
	body := struct {
		Data resources.User
	}{Data: resources.User{
		Key: resources.NewKeyInt64(request.Id, resources.USER),
		Attributes: resources.UserAttributes{
			Name:     request.Name,
			Position: request.Position,
			Surname:  request.Surname,
			Email:    request.Email,
		},
	}}

	err := c.DoRequest("PATCH", fmt.Sprintf("%s/users/%d", c.ServiceUrl, request.Id), body)
	return err
}

func (c *Connector) GetUsers(params models.GetUsersRequest) (handlers.UserListResponse, error) {
	model := handlers.UserListResponse{}

	err := c.DoRequestWithDecode(
		"GET",
		fmt.Sprintf("%s/users?%s", c.ServiceUrl, urlval.MustEncode(params)),
		nil,
		&model,
	)
	return model, err
}

func (c *Connector) GetPositions() (resources.PositionsResponse, error) {
	model := resources.PositionsResponse{}

	err := c.DoRequestWithDecode(
		http.MethodGet,
		fmt.Sprintf("%s/users/positions", c.ServiceUrl),
		nil,
		&model,
	)

	return model, err
}
