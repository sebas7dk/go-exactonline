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

// StageForDeliveryReceiptsEndpoint is responsible for communicating with
// the StageForDeliveryReceipts endpoint of the Manufacturing service.
type StageForDeliveryReceiptsEndpoint service

// StageForDeliveryReceipts:
// Service: Manufacturing
// Entity: StageForDeliveryReceipts
// URL: /api/v1/{division}/manufacturing/StageForDeliveryReceipts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingStageForDeliveryReceipts
type StageForDeliveryReceipts struct {
	// Id: ID of staged for delivery entry
	Id *types.GUID `json:"Id,omitempty"`

	// CreatedBy: ID of creating user
	CreatedBy *types.GUID `json:"CreatedBy,omitempty"`

	// CreatedByFullName: Name of the creating user
	CreatedByFullName *string `json:"CreatedByFullName,omitempty"`

	// CreatedDate: Date this Stage for delivery wa created
	CreatedDate *types.Date `json:"CreatedDate,omitempty"`

	// HasReversibleQuantity: Indicates if this StageForDeliveryReceipt has a quantity eligible to be reversed via StageForDeliveryReversals
	HasReversibleQuantity *bool `json:"HasReversibleQuantity,omitempty"`

	// IsBatch: Does the shop order receipt&#39;s item use batch numbers
	IsBatch *byte `json:"IsBatch,omitempty"`

	// IsFractionAllowedItem: Indicates if fractions (for example 0.35) are allowed for quantities of the stage for delivery receipt&#39;s item
	IsFractionAllowedItem *byte `json:"IsFractionAllowedItem,omitempty"`

	// IsSerial: Does the shop order receipt&#39;s item use serial numbers
	IsSerial *byte `json:"IsSerial,omitempty"`

	// Item: Item staged for delivery
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Code of item staged for delivery
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item staged for delivery
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemPictureUrl: Picture url of shop order item
	ItemPictureUrl *string `json:"ItemPictureUrl,omitempty"`

	// Quantity: Quantity of this StageForDeliveryReceipt
	Quantity *float64 `json:"Quantity,omitempty"`

	// RelatedId: ID of the original stage for delivery entry
	RelatedId *types.GUID `json:"RelatedId,omitempty"`

	// ShopOrder: Shop order staged for delivery
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// ShopOrderNumber: Number of shop order staged for delivery
	ShopOrderNumber *int `json:"ShopOrderNumber,omitempty"`

	// TransactionDate: Effective date of this stage for delivery receipt
	TransactionDate *types.Date `json:"TransactionDate,omitempty"`

	// Unit: Unit of measurement abbreviation of item finished
	Unit *string `json:"Unit,omitempty"`

	// UnitDescription: Unit of measurement of item finished
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// Warehouse: ID of the shop order warehouse
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of the shop order warehouse
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of the shop order warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

// List the StageForDeliveryReceipts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *StageForDeliveryReceiptsEndpoint) List(ctx context.Context, division int, all bool) ([]*StageForDeliveryReceipts, error) {
	var entities []*StageForDeliveryReceipts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/StageForDeliveryReceipts?$select=*", division)
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
