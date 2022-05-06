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
func (c *Client) Update(customerID string, request *UpdateRequest) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/%v/%v", path, customerID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusNoContent {
		return response, nil
	}
	return response, nil
}

func (c *Client) Get(customerID string) (*GetResponse, error) {
	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, customerID))
	response := &GetResponse{
		StatusResponse: resp,
	}

	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(resp.ResponseBody, &customer)

	if resp.StatusCode == http.StatusNoContent {
		return nil, err
	}

	response.Customer = &customer

	return response, err
}
