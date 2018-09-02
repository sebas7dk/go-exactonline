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

// AssemblyOrdersEndpoint is responsible for communicating with
// the AssemblyOrders endpoint of the Inventory service.
type AssemblyOrdersEndpoint service

// AssemblyOrders:
// Service: Inventory
// Entity: AssemblyOrders
// URL: /api/v1/{division}/inventory/AssemblyOrders
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=InventoryAssemblyOrders
type AssemblyOrders struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AssemblyDate: Planned date for assembly of the item
	AssemblyDate *types.Date `json:"AssemblyDate,omitempty"`

	// Description: Description of assembly order
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// FinishedQuantity: Quantity of items that have actually been assembled
	FinishedQuantity *float64 `json:"FinishedQuantity,omitempty"`

	// Item: Reference to item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode: Item code
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Notes: Notes of the assembly order
	Notes *string `json:"Notes,omitempty"`

	// OrderDate: Date of the assembly order is initiated
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderNumber: Assembly order number
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// OrderStatus: Assembly order status: 20 = Open, 30 = Partial, 50 = Complete
	OrderStatus *int `json:"OrderStatus,omitempty"`

	// PartItems: Collection of part items for assembly order
	PartItems *[]byte `json:"PartItems,omitempty"`

	// PlannedQuantity: Planned quantity of the item to be assembled
	PlannedQuantity *float64 `json:"PlannedQuantity,omitempty"`

	// StorageLocation: Reference to storage location
	StorageLocation *types.GUID `json:"StorageLocation,omitempty"`

	// StorageLocationCode: Storage location code
	StorageLocationCode *string `json:"StorageLocationCode,omitempty"`

	// StorageLocationDescription: Storage location description
	StorageLocationDescription *string `json:"StorageLocationDescription,omitempty"`

	// Warehouse: Warehouse
	Warehouse *types.GUID `json:"Warehouse,omitempty"`

	// WarehouseCode: Code of Warehouse
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription: Description of Warehouse
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`
}

// List the AssemblyOrders entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AssemblyOrdersEndpoint) List(ctx context.Context, division int, all bool) ([]*AssemblyOrders, error) {
	var entities []*AssemblyOrders
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/inventory/AssemblyOrders?$select=*", division)
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
