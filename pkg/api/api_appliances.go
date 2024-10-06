package api

import (
    "fmt"
    "net/url"
)

// Méthodes liées aux appliances

func (c *RTMSClient) GetAppliances(cloudTempleID string) ([]byte, error) {
	if cloudTempleID == "" {
		return nil, fmt.Errorf("cloudTempleID cannot be empty")
	}

	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)

	return c.doRequest("GET", "/appliances", query, nil)
}

func (c *RTMSClient) GetApplianceDetails(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}

	return c.doRequest("GET", fmt.Sprintf("/appliances/%s", id), nil, nil)
}

func (c *RTMSClient) GetApplianceServices(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}

	return c.doRequest("GET", fmt.Sprintf("/appliances/%s/services", id), nil, nil)
}

func (c *RTMSClient) SynchronizeAppliance(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}

	return c.doRequest("GET", fmt.Sprintf("/appliances/%s/synchronize", id), nil, nil)
}

func (c *RTMSClient) GetApplianceConfiguration(id, applianceVersion, pluginsPath string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}
	if applianceVersion == "" {
		return nil, fmt.Errorf("appliance version cannot be empty")
	}
	if pluginsPath == "" {
		return nil, fmt.Errorf("plugins path cannot be empty")
	}

	query := url.Values{}
	query.Set("applianceVersion", applianceVersion)
	query.Set("pluginsPath", pluginsPath)

	return c.doRequest("GET", fmt.Sprintf("/appliances/%s/configuration", id), query, nil)
}

func (c *RTMSClient) GetApplianceHealthCheck(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}

	return c.doRequest("GET", fmt.Sprintf("/appliances/%s/healthCheck", id), nil, nil)
}

func (c *RTMSClient) PostApplianceHealthCheck(id string, healthCheck map[string]interface{}) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("appliance ID cannot be empty")
	}

	return c.doRequest("POST", fmt.Sprintf("/appliances/%s/healthCheck", id), nil, healthCheck)
}