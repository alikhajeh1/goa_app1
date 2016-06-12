package client

import (
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Get prime numbers
func (c *Client) ShowNumbers(ctx context.Context, path string, lessThan int) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp3 := strconv.Itoa(lessThan)
	values.Set("lessThan", tmp3)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
