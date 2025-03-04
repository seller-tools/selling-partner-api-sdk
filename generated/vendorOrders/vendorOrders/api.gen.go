// Package vendorOrders provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package vendorOrders

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"gopkg.me/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// SubmitAcknowledgement request  with any body
	SubmitAcknowledgementWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	SubmitAcknowledgement(ctx context.Context, body SubmitAcknowledgementJSONRequestBody) (*http.Response, error)

	// GetPurchaseOrders request
	GetPurchaseOrders(ctx context.Context, params *GetPurchaseOrdersParams) (*http.Response, error)

	// GetPurchaseOrder request
	GetPurchaseOrder(ctx context.Context, purchaseOrderNumber string) (*http.Response, error)

	// GetPurchaseOrdersStatus request
	GetPurchaseOrdersStatus(ctx context.Context, params *GetPurchaseOrdersStatusParams) (*http.Response, error)
}

func (c *Client) SubmitAcknowledgementWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewSubmitAcknowledgementRequestWithBody(c.Endpoint, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) SubmitAcknowledgement(ctx context.Context, body SubmitAcknowledgementJSONRequestBody) (*http.Response, error) {
	req, err := NewSubmitAcknowledgementRequest(c.Endpoint, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetPurchaseOrders(ctx context.Context, params *GetPurchaseOrdersParams) (*http.Response, error) {
	req, err := NewGetPurchaseOrdersRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetPurchaseOrder(ctx context.Context, purchaseOrderNumber string) (*http.Response, error) {
	req, err := NewGetPurchaseOrderRequest(c.Endpoint, purchaseOrderNumber)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetPurchaseOrdersStatus(ctx context.Context, params *GetPurchaseOrdersStatusParams) (*http.Response, error) {
	req, err := NewGetPurchaseOrdersStatusRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewSubmitAcknowledgementRequest calls the generic SubmitAcknowledgement builder with application/json body
func NewSubmitAcknowledgementRequest(endpoint string, body SubmitAcknowledgementJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSubmitAcknowledgementRequestWithBody(endpoint, "application/json", bodyReader)
}

// NewSubmitAcknowledgementRequestWithBody generates requests for SubmitAcknowledgement with any type of body
func NewSubmitAcknowledgementRequestWithBody(endpoint string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/vendor/orders/v1/acknowledgements")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetPurchaseOrdersRequest generates requests for GetPurchaseOrders
func NewGetPurchaseOrdersRequest(endpoint string, params *GetPurchaseOrdersParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/vendor/orders/v1/purchaseOrders")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "limit", *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdAfter", *params.CreatedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdBefore", *params.CreatedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SortOrder != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "sortOrder", *params.SortOrder); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "nextToken", *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.IncludeDetails != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "includeDetails", *params.IncludeDetails); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ChangedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "changedAfter", *params.ChangedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ChangedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "changedBefore", *params.ChangedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PoItemState != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "poItemState", *params.PoItemState); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.IsPOChanged != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "isPOChanged", *params.IsPOChanged); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PurchaseOrderState != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "purchaseOrderState", *params.PurchaseOrderState); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.OrderingVendorCode != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "orderingVendorCode", *params.OrderingVendorCode); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPurchaseOrderRequest generates requests for GetPurchaseOrder
func NewGetPurchaseOrderRequest(endpoint string, purchaseOrderNumber string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "purchaseOrderNumber", purchaseOrderNumber)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/vendor/orders/v1/purchaseOrders/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPurchaseOrdersStatusRequest generates requests for GetPurchaseOrdersStatus
func NewGetPurchaseOrdersStatusRequest(endpoint string, params *GetPurchaseOrdersStatusParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/vendor/orders/v1/purchaseOrdersStatus")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "limit", *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SortOrder != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "sortOrder", *params.SortOrder); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "nextToken", *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdAfter", *params.CreatedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.CreatedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "createdBefore", *params.CreatedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.UpdatedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "updatedAfter", *params.UpdatedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.UpdatedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "updatedBefore", *params.UpdatedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PurchaseOrderNumber != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "purchaseOrderNumber", *params.PurchaseOrderNumber); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PurchaseOrderStatus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "purchaseOrderStatus", *params.PurchaseOrderStatus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ItemConfirmationStatus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "itemConfirmationStatus", *params.ItemConfirmationStatus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ItemReceiveStatus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "itemReceiveStatus", *params.ItemReceiveStatus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.OrderingVendorCode != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "orderingVendorCode", *params.OrderingVendorCode); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ShipToPartyId != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "shipToPartyId", *params.ShipToPartyId); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// SubmitAcknowledgement request  with any body
	SubmitAcknowledgementWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*SubmitAcknowledgementResp, error)

	SubmitAcknowledgementWithResponse(ctx context.Context, body SubmitAcknowledgementJSONRequestBody) (*SubmitAcknowledgementResp, error)

	// GetPurchaseOrders request
	GetPurchaseOrdersWithResponse(ctx context.Context, params *GetPurchaseOrdersParams) (*GetPurchaseOrdersResp, error)

	// GetPurchaseOrder request
	GetPurchaseOrderWithResponse(ctx context.Context, purchaseOrderNumber string) (*GetPurchaseOrderResp, error)

	// GetPurchaseOrdersStatus request
	GetPurchaseOrdersStatusWithResponse(ctx context.Context, params *GetPurchaseOrdersStatusParams) (*GetPurchaseOrdersStatusResp, error)
}

type SubmitAcknowledgementResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *SubmitAcknowledgementResponse
	JSON400      *SubmitAcknowledgementResponse
	JSON403      *SubmitAcknowledgementResponse
	JSON404      *SubmitAcknowledgementResponse
	JSON413      *SubmitAcknowledgementResponse
	JSON415      *SubmitAcknowledgementResponse
	JSON429      *SubmitAcknowledgementResponse
	JSON500      *SubmitAcknowledgementResponse
	JSON503      *SubmitAcknowledgementResponse
}

// Status returns HTTPResponse.Status
func (r SubmitAcknowledgementResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SubmitAcknowledgementResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPurchaseOrdersResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPurchaseOrdersResponse
	JSON400      *GetPurchaseOrdersResponse
	JSON403      *GetPurchaseOrdersResponse
	JSON404      *GetPurchaseOrdersResponse
	JSON415      *GetPurchaseOrdersResponse
	JSON429      *GetPurchaseOrdersResponse
	JSON500      *GetPurchaseOrdersResponse
	JSON503      *GetPurchaseOrdersResponse
}

// Status returns HTTPResponse.Status
func (r GetPurchaseOrdersResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPurchaseOrdersResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPurchaseOrderResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPurchaseOrderResponse
	JSON400      *GetPurchaseOrderResponse
	JSON401      *GetPurchaseOrderResponse
	JSON403      *GetPurchaseOrderResponse
	JSON404      *GetPurchaseOrderResponse
	JSON415      *GetPurchaseOrderResponse
	JSON429      *GetPurchaseOrderResponse
	JSON500      *GetPurchaseOrderResponse
	JSON503      *GetPurchaseOrderResponse
}

// Status returns HTTPResponse.Status
func (r GetPurchaseOrderResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPurchaseOrderResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPurchaseOrdersStatusResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPurchaseOrdersStatusResponse
	JSON400      *GetPurchaseOrdersStatusResponse
	JSON403      *GetPurchaseOrdersStatusResponse
	JSON404      *GetPurchaseOrdersStatusResponse
	JSON415      *GetPurchaseOrdersStatusResponse
	JSON429      *GetPurchaseOrdersStatusResponse
	JSON500      *GetPurchaseOrdersStatusResponse
	JSON503      *GetPurchaseOrdersStatusResponse
}

// Status returns HTTPResponse.Status
func (r GetPurchaseOrdersStatusResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPurchaseOrdersStatusResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// SubmitAcknowledgementWithBodyWithResponse request with arbitrary body returning *SubmitAcknowledgementResponse
func (c *ClientWithResponses) SubmitAcknowledgementWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*SubmitAcknowledgementResp, error) {
	rsp, err := c.SubmitAcknowledgementWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseSubmitAcknowledgementResp(rsp)
}

func (c *ClientWithResponses) SubmitAcknowledgementWithResponse(ctx context.Context, body SubmitAcknowledgementJSONRequestBody) (*SubmitAcknowledgementResp, error) {
	rsp, err := c.SubmitAcknowledgement(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseSubmitAcknowledgementResp(rsp)
}

// GetPurchaseOrdersWithResponse request returning *GetPurchaseOrdersResponse
func (c *ClientWithResponses) GetPurchaseOrdersWithResponse(ctx context.Context, params *GetPurchaseOrdersParams) (*GetPurchaseOrdersResp, error) {
	rsp, err := c.GetPurchaseOrders(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetPurchaseOrdersResp(rsp)
}

// GetPurchaseOrderWithResponse request returning *GetPurchaseOrderResponse
func (c *ClientWithResponses) GetPurchaseOrderWithResponse(ctx context.Context, purchaseOrderNumber string) (*GetPurchaseOrderResp, error) {
	rsp, err := c.GetPurchaseOrder(ctx, purchaseOrderNumber)
	if err != nil {
		return nil, err
	}
	return ParseGetPurchaseOrderResp(rsp)
}

// GetPurchaseOrdersStatusWithResponse request returning *GetPurchaseOrdersStatusResponse
func (c *ClientWithResponses) GetPurchaseOrdersStatusWithResponse(ctx context.Context, params *GetPurchaseOrdersStatusParams) (*GetPurchaseOrdersStatusResp, error) {
	rsp, err := c.GetPurchaseOrdersStatus(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetPurchaseOrdersStatusResp(rsp)
}

// ParseSubmitAcknowledgementResp parses an HTTP response from a SubmitAcknowledgementWithResponse call
func ParseSubmitAcknowledgementResp(rsp *http.Response) (*SubmitAcknowledgementResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &SubmitAcknowledgementResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest SubmitAcknowledgementResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetPurchaseOrdersResp parses an HTTP response from a GetPurchaseOrdersWithResponse call
func ParseGetPurchaseOrdersResp(rsp *http.Response) (*GetPurchaseOrdersResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPurchaseOrdersResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest GetPurchaseOrdersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetPurchaseOrderResp parses an HTTP response from a GetPurchaseOrderWithResponse call
func ParseGetPurchaseOrderResp(rsp *http.Response) (*GetPurchaseOrderResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPurchaseOrderResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest GetPurchaseOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetPurchaseOrdersStatusResp parses an HTTP response from a GetPurchaseOrdersStatusWithResponse call
func ParseGetPurchaseOrdersStatusResp(rsp *http.Response) (*GetPurchaseOrdersStatusResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPurchaseOrdersStatusResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest GetPurchaseOrdersStatusResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
