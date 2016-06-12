package test

import (
	"bytes"
	"fmt"
	"github.com/alikhajeh1/goa_app1/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ShowNumbersOK test setup
func ShowNumbersOK(t *testing.T, ctrl app.NumbersController) *app.GoaExampleNumbers {
	return ShowNumbersOKCtx(t, context.Background(), ctrl)
}

// ShowNumbersOKCtx test setup
func ShowNumbersOKCtx(t *testing.T, ctx context.Context, ctrl app.NumbersController) *app.GoaExampleNumbers {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/numbers"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "NumbersTest"), rw, req, prms)
	showCtx, err := app.NewShowNumbersContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.GoaExampleNumbers)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.GoaExampleNumbers", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}
