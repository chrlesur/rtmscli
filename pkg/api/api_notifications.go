package api

import (
	"net/url"
)

// Méthodes liées aux notifications

func (c *RTMSClient) GetNagiosCommands(params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/nagiosCommands", query, nil)
}

func (c *RTMSClient) GetNagiosCommandsTimePeriods(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	if cloudTempleID != "" {
		query.Set("cloudTempleId", cloudTempleID)
	}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/nagiosCommands/timePeriods", query, nil)
}

func (c *RTMSClient) ValidateNagiosPluginPackage(packageData map[string]interface{}) ([]byte, error) {
	return c.doRequest("POST", "/nagiosPlugins/validatePackage", nil, packageData)
}

func (c *RTMSClient) UpdateNagiosCommands() ([]byte, error) {
	return c.doRequest("GET", "/nagiosPlugins/updateNagiosCommands", nil, nil)
}
