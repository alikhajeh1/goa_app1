//************************************************************************//
// API "goa_app1": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/alikhajeh1/goa_app1
// --design=github.com/alikhajeh1/goa_app1/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// NumbersHref returns the resource href.
func NumbersHref() string {
	return fmt.Sprintf("/numbers")
}
