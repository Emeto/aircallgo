package aircallgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Client[E any] struct {
	config     *Config
	HTTPClient http.Client
}

type ClientResponse[E any] struct {
	Status     string
	StatusCode int
	data       E
}

func newClient() *Client[any] {
	return &Client[any]{
		config:     ParseConfig(),
		HTTPClient: http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client[any]) MakeRequest(endpoint string, method string, payload *bytes.Reader) *ClientResponse[any] {
	var response ClientResponse[any]
	// If no payload is passed (i.e. GET requests), initialize an empty Reader
	if payload == nil {
		payload = bytes.NewReader(make([]byte, 0))
	}
	req, _ := http.NewRequest(method, c.buildEndpoint(endpoint), payload)
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request: %s\n", err)
	}
	response.Status = res.Status
	if res.StatusCode == http.StatusNoContent {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal("unable to close response body")
			}
		}(res.Body)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading HTTP response: %s\n", err)
		}
		err = json.Unmarshal(body, &response.data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse JSON response: %s\n", err)
		}
	}
	return &response
}

func (c *Client[any]) buildEndpoint(uri string) string {
	return "https://" + c.config.ID + ":" + c.config.Token + "@api.aircall.io/v1/" + uri
}
