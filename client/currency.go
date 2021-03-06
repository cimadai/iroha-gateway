// API "Iroha-Gateway Server": currency Resource Client
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.1.0-dirty

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// AddCurrencyPath computes a request path to the add action of currency.
func AddCurrencyPath() string {

	return fmt.Sprintf("/currency")
}

// AddCurrency makes a request to the add action endpoint of the currency resource
func (c *Client) AddCurrency(ctx context.Context, path string, payload *CreateCurrencyRequest, contentType string) (*http.Response, error) {
	req, err := c.NewAddCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddCurrencyRequest create the request corresponding to the add action endpoint of the currency resource.
func (c *Client) NewAddCurrencyRequest(ctx context.Context, path string, payload *CreateCurrencyRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// AddValueCurrencyPath computes a request path to the addValue action of currency.
func AddValueCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s/add", param0)
}

// AddValueCurrency makes a request to the addValue action endpoint of the currency resource
func (c *Client) AddValueCurrency(ctx context.Context, path string, payload *CurrencyValueRequest, contentType string) (*http.Response, error) {
	req, err := c.NewAddValueCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddValueCurrencyRequest create the request corresponding to the addValue action endpoint of the currency resource.
func (c *Client) NewAddValueCurrencyRequest(ctx context.Context, path string, payload *CurrencyValueRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteCurrencyPath computes a request path to the delete action of currency.
func DeleteCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s", param0)
}

// DeleteCurrency makes a request to the delete action endpoint of the currency resource
func (c *Client) DeleteCurrency(ctx context.Context, path string, payload *DeleteCurrencyRequest, contentType string) (*http.Response, error) {
	req, err := c.NewDeleteCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCurrencyRequest create the request corresponding to the delete action endpoint of the currency resource.
func (c *Client) NewDeleteCurrencyRequest(ctx context.Context, path string, payload *DeleteCurrencyRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// GetAllCurrencyPath computes a request path to the getAll action of currency.
func GetAllCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s", param0)
}

// GetAllCurrency makes a request to the getAll action endpoint of the currency resource
func (c *Client) GetAllCurrency(ctx context.Context, path string, creatorPubkey string, target string, isCommitted *bool) (*http.Response, error) {
	req, err := c.NewGetAllCurrencyRequest(ctx, path, creatorPubkey, target, isCommitted)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAllCurrencyRequest create the request corresponding to the getAll action endpoint of the currency resource.
func (c *Client) NewGetAllCurrencyRequest(ctx context.Context, path string, creatorPubkey string, target string, isCommitted *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("creator_pubkey", creatorPubkey)
	values.Set("target", target)
	if isCommitted != nil {
		tmp23 := strconv.FormatBool(*isCommitted)
		values.Set("is_committed", tmp23)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SubtractValueCurrencyPath computes a request path to the subtractValue action of currency.
func SubtractValueCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s/subtract", param0)
}

// SubtractValueCurrency makes a request to the subtractValue action endpoint of the currency resource
func (c *Client) SubtractValueCurrency(ctx context.Context, path string, payload *CurrencyValueRequest, contentType string) (*http.Response, error) {
	req, err := c.NewSubtractValueCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubtractValueCurrencyRequest create the request corresponding to the subtractValue action endpoint of the currency resource.
func (c *Client) NewSubtractValueCurrencyRequest(ctx context.Context, path string, payload *CurrencyValueRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// TransferCurrencyPath computes a request path to the transfer action of currency.
func TransferCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s/transfer", param0)
}

// TransferCurrency makes a request to the transfer action endpoint of the currency resource
func (c *Client) TransferCurrency(ctx context.Context, path string, payload *CurrencyTransferRequest, contentType string) (*http.Response, error) {
	req, err := c.NewTransferCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTransferCurrencyRequest create the request corresponding to the transfer action endpoint of the currency resource.
func (c *Client) NewTransferCurrencyRequest(ctx context.Context, path string, payload *CurrencyTransferRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// UpdateCurrencyPath computes a request path to the update action of currency.
func UpdateCurrencyPath(currencyURI string) string {
	param0 := currencyURI

	return fmt.Sprintf("/currency/%s", param0)
}

// UpdateCurrency makes a request to the update action endpoint of the currency resource
func (c *Client) UpdateCurrency(ctx context.Context, path string, payload *UpdateCurrencyRequest, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateCurrencyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCurrencyRequest create the request corresponding to the update action endpoint of the currency resource.
func (c *Client) NewUpdateCurrencyRequest(ctx context.Context, path string, payload *UpdateCurrencyRequest, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
