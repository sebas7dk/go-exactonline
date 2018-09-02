// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package manufacturing

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// BillOfMaterialVersionsEndpoint is responsible for communicating with
// the BillOfMaterialVersions endpoint of the Manufacturing service.
type BillOfMaterialVersionsEndpoint service

// BillOfMaterialVersions:
// Service: Manufacturing
// Entity: BillOfMaterialVersions
// URL: /api/v1/{division}/manufacturing/BillOfMaterialVersions
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingBillOfMaterialVersions
type BillOfMaterialVersions struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// BatchQuantity: Batch Quantity of Item Version
	BatchQuantity *float64 `json:"BatchQuantity,omitempty"`

	// CadDrawingUrl: Cad drawing URL
	CadDrawingUrl *string `json:"CadDrawingUrl,omitempty"`

	// CalculatedCostPrice: Calculated Cost Price of Item Version
	CalculatedCostPrice *float64 `json:"CalculatedCostPrice,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the item version
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// IsDefault: Indicates if this is the default item version that will be assigned when a item is selected
	IsDefault *byte `json:"IsDefault,omitempty"`

	// Item: Reference to Items table
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Line notes
	Notes *string `json:"Notes,omitempty"`

	// OrderLeadDays: Order lead days for item
	OrderLeadDays *int `json:"OrderLeadDays,omitempty"`

	// ProductionLeadDays: Production lead time in days of Item version
	ProductionLeadDays *int `json:"ProductionLeadDays,omitempty"`

	// Status: Statuses of Item version: 10-Engineering change pending, 20-Engineering change approved, 30-Active &amp; 40-Historic
	Status *int `json:"Status,omitempty"`

	// StatusDescription: Description of Status
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// Type: Type of Item version: 10-Sales bill of material, 20-Manufacturing recipe
	Type *int `json:"Type,omitempty"`

	// TypeDescription: Description of Type
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// VersionDate: Version date
	VersionDate *types.Date `json:"VersionDate,omitempty"`

	// VersionNumber: Version Number
	VersionNumber *int `json:"VersionNumber,omitempty"`
}

// List the BillOfMaterialVersions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *BillOfMaterialVersionsEndpoint) List(ctx context.Context, division int, all bool) ([]*BillOfMaterialVersions, error) {
	var entities []*BillOfMaterialVersions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/BillOfMaterialVersions?$select=*", division)
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
