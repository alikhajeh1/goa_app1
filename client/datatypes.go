//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/alikhajeh1/goa_app1/design
// --out=$(GOPATH)/src/github.com/alikhajeh1/goa_app1
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "net/http"

// A list of number
type GoaExampleNumbers struct {
	// Numbers
	Numbers []int `json:"numbers,omitempty" xml:"numbers,omitempty" form:"numbers,omitempty"`
}

// DecodeGoaExampleNumbers decodes the GoaExampleNumbers instance encoded in resp body.
func (c *Client) DecodeGoaExampleNumbers(resp *http.Response) (*GoaExampleNumbers, error) {
	var decoded GoaExampleNumbers
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
