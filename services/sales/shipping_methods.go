// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package sales

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// ShippingMethodsEndpoint is responsible for communicating with
// the ShippingMethods endpoint of the Sales service.
type ShippingMethodsEndpoint service

// ShippingMethods:
// Service: Sales
// Entity: ShippingMethods
// URL: /api/v1/{division}/sales/ShippingMethods
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesShippingMethods
type ShippingMethods struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Active: Active
	Active *bool `json:"Active,omitempty"`

	// Code: Code of the shipping method
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of shipping method
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// ShippingRatesURL: Shipping method rates URL
	ShippingRatesURL *string `json:"ShippingRatesURL,omitempty"`

	// TrackingURL: Tracking URL
	TrackingURL *string `json:"TrackingURL,omitempty"`
}

func (s *ShippingMethods) GetIdentifier() types.GUID {
	return *s.ID
}

// List the ShippingMethods entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShippingMethodsEndpoint) List(ctx context.Context, division int, all bool) ([]*ShippingMethods, error) {
	var entities []*ShippingMethods
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/sales/ShippingMethods?$select=*", division)
	if err != nil {
		return nil, err
	}
	if all {
		err = s.client.ListRequestAndDoAll(ctx, u.String(), &entities)
		return entities, err
	}
	_, _, _, err = s.client.ListRequestAndDo(ctx, u.String(), &entities)
	return entities, err
}