package api

import (
	"fmt"
	"net/url"
)

// Méthodes liées aux hôtes

func (c *RTMSClient) GetHosts(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/hosts", query, nil)
}

func (c *RTMSClient) CreateHost(cloudTempleID string, hostData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/hosts", query, hostData)
}

func (c *RTMSClient) GetHostDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/hosts/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveHost(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/hosts/%s", id), nil, nil)
}

func (c *RTMSClient) UpdateHost(id string, hostData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/hosts/%s", id), nil, hostData)
}

func (c *RTMSClient) GetHostServices(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/hosts/%s/services", id), query, nil)
}

func (c *RTMSClient) UpdateHostTags(id string, tags []int) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/hosts/%s/tags", id), nil, map[string][]int{"tags": tags})
}

func (c *RTMSClient) SwitchHostMonitoring(id string, enable bool, services []int) ([]byte, error) {
	data := map[string]interface{}{
		"enable": enable,
	}
	if services != nil {
		data["services"] = services
	}
	return c.doRequest("POST", fmt.Sprintf("/hosts/%s/monitoring", id), nil, data)
}

func (c *RTMSClient) SwitchHostMonitoringNotifications(id string, enable bool, services []int) ([]byte, error) {
	data := map[string]interface{}{
		"enable": enable,
	}
	if services != nil {
		data["services"] = services
	}
	return c.doRequest("POST", fmt.Sprintf("/hosts/%s/monitoring/notifications", id), nil, data)
}

func (c *RTMSClient) GetHostsStats(cloudTempleID string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("GET", "/hosts/stats", query, nil)
}

func (c *RTMSClient) GetHostTags(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/hosts/tags", query, nil)
}

func (c *RTMSClient) CreateHostTag(cloudTempleID string, tagData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/hosts/tags", query, tagData)
}

func (c *RTMSClient) GetHostTagDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/hosts/tags/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveHostTag(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/hosts/tags/%s", id), nil, nil)
}

func (c *RTMSClient) EditHostTag(id string, tagData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/hosts/tags/%s", id), nil, tagData)
}

func (c *RTMSClient) GetHostsByTag(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/hosts/tags/%s/hosts", id), query, nil)
}
