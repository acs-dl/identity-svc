package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/requests"
	"gitlab.com/distributed_lab/acs/identity-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
	"time"
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

func (c *Connector) doRequest(method, url string, body interface{}, pagination *pgdb.OffsetPageParams) (*http.Response, error) {
	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	if pagination != nil {
		q := req.URL.Query()
		q.Add("page[limit]", strconv.FormatUint(pagination.Limit, 10))
		q.Add("page[order]", pagination.Order)
		q.Add("page[number]", strconv.FormatUint(pagination.PageNumber, 10))
		req.URL.RawQuery = q.Encode()
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	if response.StatusCode < 200 && response.StatusCode >= 300 {
		response.Body.Close()
		return response, errors.New("Bad status")
	}

	return response, nil
}

func (c *Connector) DoRequest(method, url string, body interface{}) error {
	_, err := c.doRequest(method, url, body, nil)
	return err
}

func (c *Connector) DoRequestWithDecode(method, url string, body, decodeModel interface{}, pagination *pgdb.OffsetPageParams) error {
	response, err := c.doRequest(method, url, body, pagination)
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

	err := c.DoRequestWithDecode("POST", c.ServiceUrl+"/users", body, &model, nil)
	return model, err
}

func (c *Connector) GetUser(userId int64) (resources.UserResponse, error) {
	model := resources.UserResponse{}

	err := c.DoRequestWithDecode("GET", fmt.Sprintf("%s/users/%d", c.ServiceUrl, userId), nil, &model, nil)
	return model, err
}

func (c *Connector) DeleteUser(userId int64) error {
	err := c.DoRequest("POST", fmt.Sprintf("%s/users/%d", c.ServiceUrl, userId), nil)
	return err
}

func (c *Connector) UpdateUser(request requests.UpdateUserRequest) (resources.UserResponse, error) {
	model := resources.UserResponse{}

	err := c.DoRequestWithDecode("PATCH", fmt.Sprintf("%s/users/%d", c.ServiceUrl, request.UserID), request.Data, &model, nil)
	return model, err
}

func (c *Connector) GetUsers(params pgdb.OffsetPageParams) (resources.UserListResponse, error) {
	model := resources.UserListResponse{}

	err := c.DoRequestWithDecode("GET", c.ServiceUrl+"/users", nil, &model, &params)
	return model, err
}
