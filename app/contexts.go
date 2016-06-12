//************************************************************************//
// API "goa_app1": Application Contexts
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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// ShowNumbersContext provides the numbers show action context.
type ShowNumbersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	LessThan *int
}

// NewShowNumbersContext parses the incoming request URL and body, performs validations and creates the
// context used by the numbers controller show action.
func NewShowNumbersContext(ctx context.Context, service *goa.Service) (*ShowNumbersContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowNumbersContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramLessThan := req.Params["lessThan"]
	if len(paramLessThan) > 0 {
		rawLessThan := paramLessThan[0]
		if lessThan, err2 := strconv.Atoi(rawLessThan); err2 == nil {
			tmp2 := lessThan
			tmp1 := &tmp2
			rctx.LessThan = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("lessThan", rawLessThan, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowNumbersContext) OK(r *GoaExampleNumbers) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.numbers")
	return ctx.Service.Send(ctx.Context, 200, r)
}
