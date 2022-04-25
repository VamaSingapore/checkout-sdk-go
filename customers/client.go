package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "customers"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Create a customer
func (c *Client) Create(request *CreateRequest, params *checkout.Params) (*CreateResponse, error) {
	resp, err := c.API.Post(fmt.Sprintf("/%v", path), request, params)
	response := &CreateResponse{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}

	var id string
	err = json.Unmarshal(resp.ResponseBody, &id)
	response.ID = id

	return response, err
}

// Update customer details
func (c *Client) Update(customerID string, request *Request) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/%v/%v", path, customerID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusNoContent {
		return response, err
	}
	return response, err
}
