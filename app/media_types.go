//************************************************************************//
// API "goa_app1": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/alikhajeh1/goa_app1/design
// --out=$(GOPATH)/src/github.com/alikhajeh1/goa_app1
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

// GoaExampleNumbers media type.
//
// Identifier: application/vnd.goa.example.numbers+json
type GoaExampleNumbers struct {
	// Numbers
	Numbers []int `json:"numbers,omitempty" xml:"numbers,omitempty" form:"numbers,omitempty"`
}
