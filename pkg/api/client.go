package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RTMSClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewRTMSClient(apiKey string, host string) (*RTMSClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}
	if host == "" {
		return nil, fmt.Errorf("host cannot be empty")
	}

	// Ensure the host has a scheme
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "https://" + host
	}

	// Ensure the host ends with "/v1"
	if !strings.HasSuffix(host, "/v1") {
		host = strings.TrimSuffix(host, "/") + "/v1"
	}

	return &RTMSClient{
		apiKey:  apiKey,
		baseURL: host,
		client:  &http.Client{},
	}, nil
}

func (c *RTMSClient) doRequest(method, endpoint string, query url.Values, body interface{}) ([]byte, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	if query != nil {
		u.RawQuery = query.Encode()
	}

	var reqBody []byte
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-AUTH-TOKEN", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

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

func (c *RTMSClient) GetCatalogs(cloudTempleID string, availableItems, isRoot bool) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	query.Set("availableItems", fmt.Sprintf("%t", availableItems))
	query.Set("isRoot", fmt.Sprintf("%t", isRoot))

	return c.doRequest("GET", "/catalogs", query, nil)
}

func (c *RTMSClient) GetDefaultCatalogs(availableItems, isRoot bool) ([]byte, error) {
	query := url.Values{}
	query.Set("availableItems", fmt.Sprintf("%t", availableItems))
	query.Set("isRoot", fmt.Sprintf("%t", isRoot))

	return c.doRequest("GET", "/catalogs/defaults", query, nil)
}

func (c *RTMSClient) GetCatalogItems(catalogID string, enabled *bool) ([]byte, error) {
	query := url.Values{}
	if enabled != nil {
		query.Set("enabled", fmt.Sprintf("%t", *enabled))
	}

	return c.doRequest("GET", fmt.Sprintf("/catalogs/%s/items", catalogID), query, nil)
}

func (c *RTMSClient) GetRootCatalog(catalogType string, availableItems bool) ([]byte, error) {
	query := url.Values{}
	query.Set("type", catalogType)
	query.Set("availableItems", fmt.Sprintf("%t", availableItems))

	return c.doRequest("GET", "/catalogs/root", query, nil)
}

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

func (c *RTMSClient) CheckRTMSHealth(integrationServices []int, integrationDelay int) ([]byte, error) {
	query := url.Values{}
	if len(integrationServices) > 0 {
		for _, service := range integrationServices {
			query.Add("integrationServices[]", strconv.Itoa(service))
		}
	}
	if integrationDelay > 0 {
		query.Set("integrationDelay", strconv.Itoa(integrationDelay))
	}
	return c.doRequest("GET", "/monitoring/health", query, nil)
}

func (c *RTMSClient) CheckSLACalculatorHealth(updateDelay int) ([]byte, error) {
	query := url.Values{}
	if updateDelay > 0 {
		query.Set("updateDelay", strconv.Itoa(updateDelay))
	}
	return c.doRequest("GET", "/monitoring/health/slaCalculator", query, nil)
}

func (c *RTMSClient) GetMonitoringServices(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices", query, nil)
}

func (c *RTMSClient) CreateMonitoringService(cloudTempleID string, serviceData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/monitoringServices", query, serviceData)
}

func (c *RTMSClient) GetMonitoringServiceDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveMonitoringService(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/monitoringServices/%s", id), nil, nil)
}

func (c *RTMSClient) UpdateMonitoringService(id string, serviceData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/monitoringServices/%s", id), nil, serviceData)
}

func (c *RTMSClient) GetMonitoringServiceTemplates(params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/templates", query, nil)
}

func (c *RTMSClient) GetMonitoringServicesStats(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/stats", query, nil)
}

func (c *RTMSClient) GetServiceNotifications(serviceID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/%s/notifications", serviceID), query, nil)
}

func (c *RTMSClient) GetAllNotifications(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications", query, nil)
}

func (c *RTMSClient) CreateNotification(notificationData map[string]interface{}) ([]byte, error) {
	return c.doRequest("POST", "/monitoringServices/notifications", nil, notificationData)
}

func (c *RTMSClient) GetNotificationDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/%s", id), nil, nil)
}

func (c *RTMSClient) GetTicketSuggestions(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/%s/suggest", id), nil, nil)
}

func (c *RTMSClient) AttachNotificationToTicket(id string, ticketID int) ([]byte, error) {
	data := map[string]interface{}{
		"ticket": ticketID,
	}
	return c.doRequest("POST", fmt.Sprintf("/monitoringServices/notifications/%s/attach", id), nil, data)
}

func (c *RTMSClient) DetachNotificationFromTicket(id string) ([]byte, error) {
	return c.doRequest("POST", fmt.Sprintf("/monitoringServices/notifications/%s/detach", id), nil, nil)
}

func (c *RTMSClient) GetNotificationPerimeters(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications/perimeters", query, nil)
}

func (c *RTMSClient) GetNotificationPerimeter(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/perimeters/%s", id), nil, nil)
}

func (c *RTMSClient) UpdateNotificationPerimeter(id string, perimeterData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/monitoringServices/notifications/perimeters/%s", id), nil, perimeterData)
}
