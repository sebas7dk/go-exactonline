// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package salesorder

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// SalesOrderLinesEndpoint is responsible for communicating with
// the SalesOrderLines endpoint of the SalesOrder service.
type SalesOrderLinesEndpoint service

// SalesOrderLines:
// Service: SalesOrder
// Entity: SalesOrderLines
// URL: /api/v1/{division}/salesorder/SalesOrderLines
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrderLines
type SalesOrderLines struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// CostCenter:
	CostCenter *string `json:"CostCenter,omitempty"`

	// CostCenterDescription:
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostPriceFC:
	CostPriceFC *float64 `json:"CostPriceFC,omitempty"`

	// CostUnit:
	CostUnit *string `json:"CostUnit,omitempty"`

	// CostUnitDescription:
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Item:
	Item *types.GUID `json:"Item,omitempty"`

	// ItemCode:
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription:
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemVersion:
	ItemVersion *types.GUID `json:"ItemVersion,omitempty"`

	// ItemVersionDescription:
	ItemVersionDescription *string `json:"ItemVersionDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// Margin:
	Margin *float64 `json:"Margin,omitempty"`

	// NetPrice:
	NetPrice *float64 `json:"NetPrice,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OrderID:
	OrderID *types.GUID `json:"OrderID,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// Pricelist:
	Pricelist *types.GUID `json:"Pricelist,omitempty"`

	// PricelistDescription:
	PricelistDescription *string `json:"PricelistDescription,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// PurchaseOrder:
	PurchaseOrder *types.GUID `json:"PurchaseOrder,omitempty"`

	// PurchaseOrderLine:
	PurchaseOrderLine *types.GUID `json:"PurchaseOrderLine,omitempty"`

	// PurchaseOrderLineNumber:
	PurchaseOrderLineNumber *int `json:"PurchaseOrderLineNumber,omitempty"`

	// PurchaseOrderNumber:
	PurchaseOrderNumber *int `json:"PurchaseOrderNumber,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// QuantityDelivered:
	QuantityDelivered *float64 `json:"QuantityDelivered,omitempty"`

	// QuantityInvoiced:
	QuantityInvoiced *float64 `json:"QuantityInvoiced,omitempty"`

	// ShopOrder:
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// TaxSchedule:
	TaxSchedule *types.GUID `json:"TaxSchedule,omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:"TaxScheduleCode,omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:"TaxScheduleDescription,omitempty"`

	// UnitCode:
	UnitCode *string `json:"UnitCode,omitempty"`

	// UnitDescription:
	UnitDescription *string `json:"UnitDescription,omitempty"`

	// UnitPrice:
	UnitPrice *float64 `json:"UnitPrice,omitempty"`

	// UseDropShipment:
	UseDropShipment *byte `json:"UseDropShipment,omitempty"`

	// VATAmount:
	VATAmount *float64 `json:"VATAmount,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription:
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

// List the SalesOrderLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesOrderLinesEndpoint) List(ctx context.Context, division int, all bool) ([]*SalesOrderLines, error) {
	var entities []*SalesOrderLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrderLines?$select=*", division)
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
