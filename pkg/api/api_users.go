package api

import (
	"fmt"
	"net/url"
)

// Méthodes liées aux utilisateurs

func (c *RTMSClient) GetUsers(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/users", query, nil)
}

func (c *RTMSClient) CreateUser(cloudTempleID string, userData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/users", query, userData)
}

func (c *RTMSClient) GetUserDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/users/%s", id), nil, nil)
}

func (c *RTMSClient) UpdateUser(id string, userData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/users/%s", id), nil, userData)
}

func (c *RTMSClient) GetWhoAmI() ([]byte, error) {
	return c.doRequest("GET", "/users/whoami", nil, nil)
}

func (c *RTMSClient) GetNotAssignedUser() ([]byte, error) {
	return c.doRequest("GET", "/users/notAssigned", nil, nil)
}

func (c *RTMSClient) GetOnDelegationUser() ([]byte, error) {
	return c.doRequest("GET", "/users/onDelegation", nil, nil)
}
