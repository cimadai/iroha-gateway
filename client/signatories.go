// API "Iroha-Gateway Server": signatories Resource Client
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

// AddSignatoriesPath computes a request path to the add action of signatories.
func AddSignatoriesPath(target string) string {
	param0 := target

	return fmt.Sprintf("/accounts/%s/signatories", param0)
}

// AddSignatories makes a request to the add action endpoint of the signatories resource
func (c *Client) AddSignatories(ctx context.Context, path string, payload *SignatoryRequest, contentType string) (*http.Response, error) {
	req, err := c.NewAddSignatoriesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddSignatoriesRequest create the request corresponding to the add action endpoint of the signatories resource.
func (c *Client) NewAddSignatoriesRequest(ctx context.Context, path string, payload *SignatoryRequest, contentType string) (*http.Request, error) {
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

// DeleteSignatoriesPath computes a request path to the delete action of signatories.
func DeleteSignatoriesPath(target string, sig string) string {
	param0 := target
	param1 := sig

	return fmt.Sprintf("/accounts/%s/signatories/%s", param0, param1)
}

// DeleteSignatories makes a request to the delete action endpoint of the signatories resource
func (c *Client) DeleteSignatories(ctx context.Context, path string, payload *DeleteSignatoryRequest, contentType string) (*http.Response, error) {
	req, err := c.NewDeleteSignatoriesRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteSignatoriesRequest create the request corresponding to the delete action endpoint of the signatories resource.
func (c *Client) NewDeleteSignatoriesRequest(ctx context.Context, path string, payload *DeleteSignatoryRequest, contentType string) (*http.Request, error) {
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

// GetAllSignatoriesPath computes a request path to the getAll action of signatories.
func GetAllSignatoriesPath(target string) string {
	param0 := target

	return fmt.Sprintf("/accounts/%s/signatories", param0)
}

// GetAllSignatories makes a request to the getAll action endpoint of the signatories resource
func (c *Client) GetAllSignatories(ctx context.Context, path string, creatorPubkey *string, isCommitted *bool) (*http.Response, error) {
	req, err := c.NewGetAllSignatoriesRequest(ctx, path, creatorPubkey, isCommitted)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAllSignatoriesRequest create the request corresponding to the getAll action endpoint of the signatories resource.
func (c *Client) NewGetAllSignatoriesRequest(ctx context.Context, path string, creatorPubkey *string, isCommitted *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if creatorPubkey != nil {
		values.Set("creator_pubkey", *creatorPubkey)
	}
	if isCommitted != nil {
		tmp24 := strconv.FormatBool(*isCommitted)
		values.Set("is_committed", tmp24)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
