// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package logistics

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// UnitsEndpoint is responsible for communicating with
// the Units endpoint of the Logistics service.
type UnitsEndpoint service

// Units:
// Service: Logistics
// Entity: Units
// URL: /api/v1/{division}/logistics/Units
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=LogisticsUnits
type Units struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Active: Indicates whether a unit is in use
	Active *bool `json:"Active,omitempty"`

	// Code: Unique code for the unit
	Code *string `json:"Code,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Main: Indicates the main unit per division. Will be used when creating new item
	Main *byte `json:"Main,omitempty"`

	// TimeUnit: If Type = &#39;T&#39; (time) then this fields indicates the type of time frame. yy = Year, mm = Month, wk = Week, dd = Day, hh = Hour, mi = Minute, ss = Second
	TimeUnit *string `json:"TimeUnit,omitempty"`

	// Type: Type of unit. Type &#39;Time&#39; is especially important for contracts.
	Type *string `json:"Type,omitempty"`
}

// List the Units entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *UnitsEndpoint) List(ctx context.Context, division int, all bool) ([]*Units, error) {
	var entities []*Units
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/logistics/Units?$select=*", division)
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
