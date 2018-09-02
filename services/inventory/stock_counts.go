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

// StockCountsEndpoint is responsible for communicating with
// the StockCounts endpoint of the Inventory service.
type StockCountsEndpoint service

// StockCounts:
// Service: Inventory
// Entity: StockCounts
// URL: /api/v1/{division}/inventory/StockCounts
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryStockCounts
type StockCounts struct {
	// StockCountID: Primary key
	StockCountID *types.GUID `json:"StockCountID,omitempty"`

	// CountedBy: Stock count user
	CountedBy *types.GUID `json:"CountedBy,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the stock count
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EntryNumber: Entry number of the stock transaction
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OffsetGLInventory: Offset GL account of inventory
	OffsetGLInventory *types.GUID `json:"OffsetGLInventory,omitempty"`

	// OffsetGLInventoryCode: GLAccount code
	OffsetGLInventoryCode *string `json:"OffsetGLInventoryCode,omitempty"`

	// OffsetGLInventoryDescription: GLAccount description
	OffsetGLInventoryDescription *string `json:"OffsetGLInventoryDescription,omitempty"`

	// Source: Source of stock count entry: 1-Manual entry, 2-Import, 3-Stock count, 4-Web service
	Source *int `json:"Source,omitempty"`

	// Status: Stock count status: 12-Draft, 21-Processed
	Status *int `json:"Status,omitempty"`

	// StockCountDate: Stock count date
	StockCountDate *types.Date `json:"StockCountDate,omitempty"`

	// StockCountLines: Collection of stock count lines
	StockCountLines *[]byte `json:"StockCountLines,omitempty"`

	// StockCountNumber: Human readable id of the stock count
	StockCountNumber *int `json:"StockCountNumber,omitempty"`

	// Warehouse: Warehouse
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of Warehouse
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of Warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

// List the StockCounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *StockCountsEndpoint) List(ctx context.Context, division int, all bool) ([]*StockCounts, error) {
	var entities []*StockCounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/StockCounts?$select=*", division)
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
