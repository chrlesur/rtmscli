package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// Méthodes liées aux tickets, pièces jointes, commentaires et tags

func (c *RTMSClient) GetTickets(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	if cloudTempleID != "" {
		query.Set("cloudTempleId", cloudTempleID)
	}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/tickets", query, nil)
}

func (c *RTMSClient) CreateTicket(cloudTempleID string, ticketData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/tickets", query, ticketData)
}

func (c *RTMSClient) GetTicketsCount(cloudTempleID string, status int) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	if status >= 0 {
		query.Set("status", fmt.Sprintf("%d", status))
	}
	return c.doRequest("GET", "/tickets/count", query, nil)
}

func (c *RTMSClient) GetTicketDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tickets/%s", id), nil, nil)
}

func (c *RTMSClient) EditTicket(id string, ticketData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tickets/%s", id), nil, ticketData)
}

func (c *RTMSClient) GetTicketCatalogs(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/tickets/%s/catalogs", id), query, nil)
}

func (c *RTMSClient) GetTicketsStats(cloudTempleID string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("GET", "/tickets/stats", query, nil)
}

func (c *RTMSClient) ListTicketAttachments(ticketID string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tickets/%s/attachments", ticketID), nil, nil)
}

func (c *RTMSClient) UploadTicketAttachment(ticketID string, filename string, content []byte) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("attachment", filename)
	if err != nil {
		return nil, err
	}
	_, err = part.Write(content)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+fmt.Sprintf("/tickets/%s/attachments", ticketID), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-AUTH-TOKEN", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *RTMSClient) DownloadTicketAttachment(attachmentID string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tickets/attachments/%s", attachmentID), nil, nil)
}

func (c *RTMSClient) RemoveTicketAttachment(attachmentID string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/tickets/attachments/%s", attachmentID), nil, nil)
}

func (c *RTMSClient) GetTicketComments(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/tickets/comments", query, nil)
}

func (c *RTMSClient) GetTicketCommentsByTicket(ticketID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/tickets/%s/comments", ticketID), query, nil)
}

func (c *RTMSClient) PostTicketComment(ticketID string, commentData map[string]interface{}) ([]byte, error) {
	return c.doRequest("POST", fmt.Sprintf("/tickets/%s/comments", ticketID), nil, commentData)
}

func (c *RTMSClient) EditTicketComment(commentID string, commentData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tickets/comments/%s", commentID), nil, commentData)
}

func (c *RTMSClient) GetTicketTags(cloudTempleID string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", "/tickets/tags", query, nil)
}

func (c *RTMSClient) CreateTicketTag(cloudTempleID string, tagData map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	query.Set("cloudTempleId", cloudTempleID)
	return c.doRequest("POST", "/tickets/tags", query, tagData)
}

func (c *RTMSClient) GetTicketTagDetails(id string) ([]byte, error) {
	return c.doRequest("GET", fmt.Sprintf("/tickets/tags/%s", id), nil, nil)
}

func (c *RTMSClient) RemoveTicketTag(id string) ([]byte, error) {
	return c.doRequest("DELETE", fmt.Sprintf("/tickets/tags/%s", id), nil, nil)
}

func (c *RTMSClient) EditTicketTag(id string, tagData map[string]interface{}) ([]byte, error) {
	return c.doRequest("PATCH", fmt.Sprintf("/tickets/tags/%s", id), nil, tagData)
}

func (c *RTMSClient) GetTicketsByTag(id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/tickets/tags/%s/tickets", id), query, nil)
}
