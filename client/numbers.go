package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// ShowNumbersPath computes a request path to the show action of numbers.
func ShowNumbersPath() string {
	return fmt.Sprintf("/numbers")
}

// Get prime numbers
func (c *Client) ShowNumbers(ctx context.Context, path string, lessThan *int) (*http.Response, error) {
	req, err := c.NewShowNumbersRequest(ctx, path, lessThan)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowNumbersRequest create the request corresponding to the show action endpoint of the numbers resource.
func (c *Client) NewShowNumbersRequest(ctx context.Context, path string, lessThan *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if lessThan != nil {
		tmp2 := strconv.Itoa(*lessThan)
		values.Set("lessThan", tmp2)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
