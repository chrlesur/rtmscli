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
	apiKey       string
	baseURL      string
	client       *http.Client
	isBase64Func func(string) bool
	debug        bool // New debug field
}

func NewRTMSClient(apiKey string, host string, isBase64Func func(string) bool) (*RTMSClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}
	if host == "" {
		return nil, fmt.Errorf("host cannot be empty")
	}

	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "https://" + host
	}

	if !strings.HasSuffix(host, "/v1") {
		host = strings.TrimSuffix(host, "/") + "/v1"
	}

	return &RTMSClient{
		apiKey:       apiKey,
		baseURL:      host,
		client:       &http.Client{},
		isBase64Func: isBase64Func,
	}, nil
}

func (c *RTMSClient) SetDebug(debug bool) {
	c.debug = debug
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

	if c.debug {
		fmt.Printf("Request: %s %s\n", method, u.String())
		if reqBody != nil {
			fmt.Printf("Request Body: %s\n", string(reqBody))
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if c.debug {
		fmt.Printf("Response Status: %d\n", resp.StatusCode)
		fmt.Printf("Response Body: %s\n", string(respBody))
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *RTMSClient) GetViewItems(viewType, id string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	return c.doRequest("GET", fmt.Sprintf("/views/%s/%s", viewType, id), query, nil)
}
func (c *RTMSClient) StreamData(endpoint string, params map[string]string, batchSize int) (<-chan interface{}, <-chan error) {
	dataChan := make(chan interface{})
	errChan := make(chan error, 1)

	go func() {
		defer close(dataChan)
		defer close(errChan)

		offset := 0
		for {
			// Copie les paramètres originaux
			queryParams := make(url.Values)
			for k, v := range params {
				queryParams.Set(k, v)
			}

			// Ajoute les paramètres de pagination
			queryParams.Set("page", strconv.Itoa(offset/batchSize+1))
			queryParams.Set("itemsPerPage", strconv.Itoa(batchSize))

			// Effectue la requête
			response, err := c.doRequest("GET", endpoint, queryParams, nil)
			if err != nil {
				errChan <- fmt.Errorf("erreur lors de la requête API : %w", err)
				return
			}

			// Parse la réponse JSON
			var paginatedResp struct {
				Data       []interface{} `json:"data"`
				Pagination struct {
					Total int `json:"total"`
				} `json:"pagination"`
			}
			err = json.Unmarshal(response, &paginatedResp)
			if err != nil {
				errChan <- fmt.Errorf("erreur lors du décodage JSON : %w", err)
				return
			}

			// Envoie les données dans le canal
			for _, item := range paginatedResp.Data {
				dataChan <- item
			}

			// Vérifie si on a atteint la fin des données
			if offset+len(paginatedResp.Data) >= paginatedResp.Pagination.Total {
				return
			}

			// Met à jour l'offset pour la prochaine requête
			offset += len(paginatedResp.Data)
		}
	}()

	return dataChan, errChan
}
