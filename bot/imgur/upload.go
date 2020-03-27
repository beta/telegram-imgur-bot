package imgur

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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

	// Send request.
	respBody, err := c.request(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[UploadImage] %v", err)
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

// Image represents an uploaded image on Imgur.
type Image struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	URL        string `json:"link"`
	DeleteHash string `json:"deletehash"`
}
