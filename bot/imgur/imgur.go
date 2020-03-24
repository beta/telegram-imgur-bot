package imgur

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

// Image represents an uploaded image on Imgur.
type Image struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	URL        string `json:"link"`
	DeleteHash string `json:"deletehash"`
}

// UploadImage uploads an image to Imgur.
func (c *Client) UploadImage(ctx context.Context, imageURL, title string) (*Image, error) {
	const endpoint = "https://api.imgur.com/3/image"

	form := url.Values{}

	form.Set("image", imageURL)
	form.Set("type", "url")
	form.Set("title", title)

	// Create an HTTP request.
	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("[UploadImage] error while creating HTTP request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", c.clientID))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

	// Send the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[UploadImage] error while sending API request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[UploadImage] unsuccessful HTTP status code returned by Imgur API: %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("[UploadImage] nil body returned by Imgur API")
	}

	// Parse the response.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[UploadImage] error while reading the response body: %v", err)
	}
	log.Printf("[UploadImage] response body from Imgur: %s", string(respBody))

	ret := new(struct {
		Image   *Image `json:"data"`
		Success bool   `json:"success"`
		Status  int    `json:"status"`
	})
	if err := json.Unmarshal(respBody, ret); err != nil {
		return nil, fmt.Errorf("[UploadImage] error while unmarshaling the response body: %v", err)
	}
	if !ret.Success || ret.Status != http.StatusOK {
		return nil, fmt.Errorf("[UploadImage] unsuccessful status code returned by Imgur API: %d", ret.Status)
	}

	return ret.Image, nil
}
