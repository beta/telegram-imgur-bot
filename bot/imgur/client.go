package imgur

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Init initializes Imgur's OAuth2 client.
func Init(clientID string) {
	initOnce.Do(func() {
		client = &Client{clientID: clientID}
	})
}

// GetClient returns the initialized client instance.
func GetClient() *Client {
	return client
}

// Client is an API client for Imgur.
type Client struct {
	clientID string
}

var (
	client   *Client
	initOnce = new(sync.Once)
)

func (c *Client) request(ctx context.Context, req *http.Request) ([]byte, error) {
	// Set headers.
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", c.clientID))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

	// Send the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending API request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessful HTTP status code returned by Imgur API: %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("nil body returned by Imgur API")
	}

	// Parse the response.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response body: %v", err)
	}

	return respBody, nil
}
