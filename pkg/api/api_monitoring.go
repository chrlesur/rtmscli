package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// Méthodes liées au monitoring et aux services de monitoring

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

func (c *RTMSClient) GetNotificationStaffs(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications/staffs", query, nil)
}

func (c *RTMSClient) GetNotificationStaff(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/staffs/%s", id), nil, nil)
}

func (c *RTMSClient) GetNotificationTimePeriods(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications/timePeriods", query, nil)
}

func (c *RTMSClient) GetNotificationTimePeriodStops(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications/timePeriodStops", query, nil)
}

func (c *RTMSClient) CreateNotificationTimePeriodStop(cloudTempleID string, stopData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/monitoringServices/notifications/timePeriodStops", query, stopData)
}

func (c *RTMSClient) GetNotificationTimePeriodStop(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/timePeriodStops/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveNotificationTimePeriodStop(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/monitoringServices/notifications/timePeriodStops/%s", id), nil, nil)
}

func (c *RTMSClient) GetNotificationTriggers(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/monitoringServices/notifications/triggers", query, nil)
}

func (c *RTMSClient) GetNotificationTriggerDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/notifications/triggers/%s", id), nil, nil)
}

func (c *RTMSClient) GetMetricHistory(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/%s/metricHistory", id), query, nil)
}

func (c *RTMSClient) GetGraphConfigurations(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/monitoringServices/%s/graphs", id), query, nil)
}
