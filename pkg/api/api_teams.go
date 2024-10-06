package api

import (
	"fmt"
	"net/url"
)

// Méthodes liées aux équipes

func (c *RTMSClient) GetTeams(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/teams", query, nil)
}

func (c *RTMSClient) CreateTeam(cloudTempleID string, teamData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/teams", query, teamData)
}

func (c *RTMSClient) GetDefaultTeams(params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/teams/defaults", query, nil)
}

func (c *RTMSClient) GetTeamDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/teams/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveTeam(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/teams/%s", id), nil, nil)
}

func (c *RTMSClient) EditTeam(id string, teamData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/teams/%s", id), nil, teamData)
}
