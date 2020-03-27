package imgur

import (
	"context"
	"fmt"
	"net/http"
)

// DeleteImage deletes an image from Imgur.
func (c *Client) DeleteImage(ctx context.Context, deleteHash string) error {
	const endpoint = "https://api.imgur.com/3/image/"
	url := endpoint + deleteHash

	// Create an HTTP request.
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("[DeleteImage] error while creating HTTP request: %v", err)
	}

	// Send request.
	_, err = c.request(ctx, req)
	if err != nil {
		return fmt.Errorf("[DeleteImage] %v", err)
	}

	return nil
}
