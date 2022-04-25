package customers

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request -
	Request struct {
		*Customer
	}

	// CreateRequest -
	CreateRequest struct {
		Email    string            `json:"email"`
		Name     string            `json:"name,omitempty"`
		Phone    *common.Phone     `json:"phone,omitempty"`
		Metadata map[string]string `json:"metadata,omitempty"`
	}

	// Customer -
	Customer struct {
		Email   string `json:"email,omitempty"`
		Name    string `json:"name,omitempty"`
		Default string `json:"default,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	}

	// CreateResponse -
	CreateResponse struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		ID             string                   `json:"id"`
	}
)
