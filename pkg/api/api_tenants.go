package api

import (
	"fmt"
	"net/url"
)

// Méthodes liées aux tenants

func (c *RTMSClient) GetTenants(params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/tenants", query, nil)
}

func (c *RTMSClient) CreateTenant(tenantData map[string]interface{}) ([]byte, error) {
	return c.doRequest("POST", "/tenants", nil, tenantData)
}

func (c *RTMSClient) GetTenantDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tenants/%s", id), nil, nil)
}

func (c *RTMSClient) GetTenantContacts(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tenants/%s/contacts", id), nil, nil)
}

func (c *RTMSClient) RequestTenantDeletion(id string, delete bool) ([]byte, error) {
	data := map[string]interface{}{
		"delete": delete,
	}
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/deletionRequest", id), nil, data)
}

func (c *RTMSClient) GetTenantSSHKeys(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tenants/%s/sshKeys", id), nil, nil)
}

func (c *RTMSClient) GenerateTenantSSHKey(id string, keyData map[string]interface{}) ([]byte, error) {
	return c.doRequest("POST", fmt.Sprintf("/tenants/%s/sshKeys", id), nil, keyData)
}

func (c *RTMSClient) DeleteTenantSSHKey(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/tenants/sshKeys/%s", id), nil, nil)
}

func (c *RTMSClient) UpdateTenantSSHKey(id string, keyData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/sshKeys/%s", id), nil, keyData)
}

func (c *RTMSClient) GetTenantWorkflowEmails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tenants/%s/workflowEmails", id), nil, nil)
}

func (c *RTMSClient) EditTenantWorkflowEmailsGeneralities(id string, data map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/workflowEmails/generalities", id), nil, data)
}

func (c *RTMSClient) EditTenantWorkflowEmailsCreateTicket(id string, data map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/workflowEmails/createTicket", id), nil, data)
}

func (c *RTMSClient) EditTenantWorkflowEmailsUpdateTicket(id string, data map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/workflowEmails/updateTicket", id), nil, data)
}

func (c *RTMSClient) EditTenantWorkflowEmailsValidationClientTicket(id string, data map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/workflowEmails/validationClientTicket", id), nil, data)
}

func (c *RTMSClient) EditTenantWorkflowEmailsCloseTicket(id string, data map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tenants/%s/workflowEmails/closeTicket", id), nil, data)
}
