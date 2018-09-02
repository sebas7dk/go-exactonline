// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package inventory

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// StockCountLinesEndpoint is responsible for communicating with
// the StockCountLines endpoint of the Inventory service.
type StockCountLinesEndpoint service

// StockCountLines:
// Service: Inventory
// Entity: StockCountLines
// URL: /api/v1/{division}/inventory/StockCountLines
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryStockCountLines
type StockCountLines struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// BatchNumbers: The collection of batch numbers that belong to the items included in this stock count
	BatchNumbers *[]byte `json:"BatchNumbers,omitempty"`

	// CostPrice: Cost price of the item that is used to create the stock count
	CostPrice *float64 `json:"CostPrice,omitempty"`

	// CountedBy: Counted by
	CountedBy *types.GUID `json:"CountedBy,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Item: Reference to the item for which the stock is counted
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemCostPrice: Current standard/actual item cost price
	ItemCostPrice *float64 `json:"ItemCostPrice,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemDivisable: Indicates if fractional quantities of the item can be used, for example quantity = 0.4
	ItemDivisable *bool `json:"ItemDivisable,omitempty"`

	// LineNumber: Line number
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// QuantityDifference: The difference between the current quantity in stock and the new quantity in stock. For example specify -1 for this field to correct the quantity if one item in stock is broken.
	QuantityDifference *float64 `json:"QuantityDifference,omitempty"`

	// QuantityInStock: The current quantity available in stock
	QuantityInStock *float64 `json:"QuantityInStock,omitempty"`

	// QuantityNew: The new quantity in stock. Use this field to correct the quantity when the items in stock are physically counted.
	QuantityNew *float64 `json:"QuantityNew,omitempty"`

	// SerialNumbers: The collection of serial numbers that belong to the items included in this stock count
	SerialNumbers *[]byte `json:"SerialNumbers,omitempty"`

	// StockCountID: Identifies the stock count. All the lines of a stock count have the same StockCountID
	StockCountID *types.GUID `json:"StockCountID,omitempty"`

	// StockKeepingUnit: Stock item&#39;s unit description
	StockKeepingUnit *string `json:"StockKeepingUnit,omitempty"`

	// StorageLocation: This property is package specific (Stock count can have multiple lines for the same item only if it is for multiple storage locations).
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Storage location code
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Storage location description
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`
}

// List the StockCountLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *StockCountLinesEndpoint) List(ctx context.Context, division int, all bool) ([]*StockCountLines, error) {
	var entities []*StockCountLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockCountLines?$select=*", division)
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
