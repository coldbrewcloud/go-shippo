package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/corvallis3d/go-shippo/errors"
	"github.com/corvallis3d/go-shippo/models"
)

const (
	shippoAPIBaseURL = "https://api.goshippo.com/v1"
)

type Client struct {
	privateToken string
	apiVersion   string
	logger       *log.Logger
}

type listOutputCallback func(v json.RawMessage) error

// NewClient creates a new Shippo API client instance.
func NewClient(privateToken, apiVersion string) *Client {
	return &Client{
		privateToken: privateToken,
		apiVersion:   apiVersion,
	}
}

// SetTraceLogger sets a new trace logger and returns the old logger.
// If logger is not nil, Client will output all internal messages to the logger.
func (c *Client) SetTraceLogger(logger *log.Logger) *log.Logger {
	oldLogger := c.logger
	c.logger = logger
	return oldLogger
}

func (c *Client) do(method, path string, input, output interface{}) error {
	url := shippoAPIBaseURL + path

	req, err := c.createRequest(method, url, input)
	if err != nil {
		return fmt.Errorf("Error creating request object: %s", err.Error())
	}

	if err := c.executeRequest(req, output); err != nil {
		if aerr, ok := err.(*errors.APIError); ok {
			return aerr
		}
		return fmt.Errorf("Error executing request: %s", err.Error())
	}

	return nil
}

func (c *Client) doList(method, path string, input interface{}, outputCallback listOutputCallback) error {
	nextURL := shippoAPIBaseURL + path + "?results=25"

	for {
		req, err := c.createRequest(method, nextURL, input)
		if err != nil {
			return fmt.Errorf("Error creating request object: %s", err.Error())
		}

		listOutput := &models.ListAPIOutput{}
		if err := c.executeRequest(req, listOutput); err != nil {
			if aerr, ok := err.(*errors.APIError); ok {
				return aerr
			}
			return fmt.Errorf("Error executing request: %s", err.Error())
		}

		for _, v := range listOutput.Results {
			if err := outputCallback(v); err != nil {
				return fmt.Errorf("Error unmarshalling output item: %s", err.Error())
			}
		}

		if listOutput.NextPageURL == nil {
			break
		}

		nextURL = *listOutput.NextPageURL
	}

	return nil
}

func (c *Client) createRequest(method, url string, bodyObject interface{}) (req *http.Request, err error) {
	var reqBodyDebug []byte

	if c.logger != nil {
		defer func() {
			if err != nil {
				c.logPrintf("Client.createRequest() error: %s", err.Error())
				return
			} else if req == nil {
				c.logPrintf("Client.createRequest() req=nil, error=nil")
				return
			}

			headers := []string{}
			for hk, hva := range req.Header {
				for _, hv := range hva {
					headers = append(headers, fmt.Sprintf("%s=%s", hk, hv))
				}
			}

			body := ""
			if reqBodyDebug != nil {
				body = string(reqBodyDebug)
			}

			c.logPrintf("Client.createRequest() HTTP request created: method=%q, url=%q, headers=%q, body=%q",
				req.Method, req.URL.String(), strings.Join(headers, ","), body)
		}()
	}

	var reqBody io.Reader
	if bodyObject != nil {
		data, err := json.Marshal(bodyObject)
		if err != nil {
			return nil, fmt.Errorf("Error marshaling body object: %s", err.Error())
		}

		reqBodyDebug = data

		reqBody = bytes.NewBuffer(data)
	}

	req, err = http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "ShippoToken "+c.privateToken)
	if c.apiVersion != "" {
		req.Header.Set("Shippo-API-Version", c.apiVersion)
	}

	// no keep-alive
	req.Header.Set("Connection", "close")
	req.Close = true

	return req, nil
}

func (c *Client) executeRequest(req *http.Request, output interface{}) (err error) {
	if c.logger != nil {
		defer func() {
			if err != nil {
				c.logPrintf("Client.executeRequest() error: %s", err.Error())
			}
		}()
	}

	httpClient := http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %s", err.Error())
	}
	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body data: %s", err.Error())
	}

	if c.logger != nil {
		c.logPrintf("Client.executeRequest() response: status=%q, body=%q", res.Status, string(resData))
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if output != nil && len(resData) > 0 {
			if err := json.Unmarshal(resData, output); err != nil {
				return fmt.Errorf("Error unmarshaling response data: %s", err.Error())
			}
		}

		return nil
	}

	return &errors.APIError{
		Status:       res.StatusCode,
		ResponseBody: resData,
	}
}

func (c *Client) logPrintf(format string, args ...interface{}) {
	if c.logger != nil {
		c.logger.Printf(format, args...)
	}
}
