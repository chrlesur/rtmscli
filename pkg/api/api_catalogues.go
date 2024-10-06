package api

import (
	"fmt"
	"net/url"
)

// Méthodes liées aux catalogues

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
