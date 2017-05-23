// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "Iroha-Gateway Server": transactions TestHelpers
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.2.0

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/soramitsu/iroha-gateway/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// GetAllTransactionsBadRequest runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllTransactionsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TransactionsController, currencyURI string, creatorPubkey string, isCommitted *bool, target string) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		query["target"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/transactions/%v", currencyURI),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["currency_uri"] = []string{fmt.Sprintf("%v", currencyURI)}
	{
		sliceVal := []string{creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		prms["target"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TransactionsTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllTransactionsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Message)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetAllTransactionsInternalServerError runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllTransactionsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TransactionsController, currencyURI string, creatorPubkey string, isCommitted *bool, target string) (http.ResponseWriter, *app.Message) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		query["target"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/transactions/%v", currencyURI),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["currency_uri"] = []string{fmt.Sprintf("%v", currencyURI)}
	{
		sliceVal := []string{creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		prms["target"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TransactionsTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllTransactionsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Message)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Message", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetAllTransactionsOK runs the method GetAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetAllTransactionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.TransactionsController, currencyURI string, creatorPubkey string, isCommitted *bool, target string) (http.ResponseWriter, *app.Transactions) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{creatorPubkey}
		query["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		query["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		query["target"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/transactions/%v", currencyURI),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["currency_uri"] = []string{fmt.Sprintf("%v", currencyURI)}
	{
		sliceVal := []string{creatorPubkey}
		prms["creator_pubkey"] = sliceVal
	}
	if isCommitted != nil {
		sliceVal := []string{fmt.Sprintf("%v", *isCommitted)}
		prms["is_committed"] = sliceVal
	}
	{
		sliceVal := []string{target}
		prms["target"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TransactionsTest"), rw, req, prms)
	getAllCtx, _err := app.NewGetAllTransactionsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GetAll(getAllCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Transactions
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Transactions)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Transactions", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
