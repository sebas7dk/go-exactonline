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

// SalesOrdersEndpoint is responsible for communicating with
// the SalesOrders endpoint of the SalesOrder service.
type SalesOrdersEndpoint service

// SalesOrders:
// Service: SalesOrder
// Entity: SalesOrders
// URL: /api/v1/{division}/salesorder/SalesOrders
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=SalesOrderSalesOrders
type SalesOrders struct {
	// OrderID:
	OrderID *types.GUID `json:"OrderID,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountDiscount:
	AmountDiscount *float64 `json:"AmountDiscount,omitempty"`

	// AmountDiscountExclVat:
	AmountDiscountExclVat *float64 `json:"AmountDiscountExclVat,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// AmountFCExclVat:
	AmountFCExclVat *float64 `json:"AmountFCExclVat,omitempty"`

	// ApprovalStatus:
	ApprovalStatus *int `json:"ApprovalStatus,omitempty"`

	// ApprovalStatusDescription:
	ApprovalStatusDescription *string `json:"ApprovalStatusDescription,omitempty"`

	// Approved:
	Approved *types.Date `json:"Approved,omitempty"`

	// Approver:
	Approver *types.GUID `json:"Approver,omitempty"`

	// ApproverFullName:
	ApproverFullName *string `json:"ApproverFullName,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// DeliverTo:
	DeliverTo *types.GUID `json:"DeliverTo,omitempty"`

	// DeliverToContactPerson:
	DeliverToContactPerson *types.GUID `json:"DeliverToContactPerson,omitempty"`

	// DeliverToContactPersonFullName:
	DeliverToContactPersonFullName *string `json:"DeliverToContactPersonFullName,omitempty"`

	// DeliverToName:
	DeliverToName *string `json:"DeliverToName,omitempty"`

	// DeliveryAddress:
	DeliveryAddress *types.GUID `json:"DeliveryAddress,omitempty"`

	// DeliveryDate:
	DeliveryDate *types.Date `json:"DeliveryDate,omitempty"`

	// DeliveryStatus:
	DeliveryStatus *int `json:"DeliveryStatus,omitempty"`

	// DeliveryStatusDescription:
	DeliveryStatusDescription *string `json:"DeliveryStatusDescription,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Discount:
	Discount *float64 `json:"Discount,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentNumber:
	DocumentNumber *int `json:"DocumentNumber,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// InvoiceStatus:
	InvoiceStatus *int `json:"InvoiceStatus,omitempty"`

	// InvoiceStatusDescription:
	InvoiceStatusDescription *string `json:"InvoiceStatusDescription,omitempty"`

	// InvoiceTo:
	InvoiceTo *types.GUID `json:"InvoiceTo,omitempty"`

	// InvoiceToContactPerson:
	InvoiceToContactPerson *types.GUID `json:"InvoiceToContactPerson,omitempty"`

	// InvoiceToContactPersonFullName:
	InvoiceToContactPersonFullName *string `json:"InvoiceToContactPersonFullName,omitempty"`

	// InvoiceToName:
	InvoiceToName *string `json:"InvoiceToName,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// OrderDate:
	OrderDate *types.Date `json:"OrderDate,omitempty"`

	// OrderedBy:
	OrderedBy *types.GUID `json:"OrderedBy,omitempty"`

	// OrderedByContactPerson:
	OrderedByContactPerson *types.GUID `json:"OrderedByContactPerson,omitempty"`

	// OrderedByContactPersonFullName:
	OrderedByContactPersonFullName *string `json:"OrderedByContactPersonFullName,omitempty"`

	// OrderedByName:
	OrderedByName *string `json:"OrderedByName,omitempty"`

	// OrderNumber:
	OrderNumber *int `json:"OrderNumber,omitempty"`

	// PaymentCondition:
	PaymentCondition *string `json:"PaymentCondition,omitempty"`

	// PaymentConditionDescription:
	PaymentConditionDescription *string `json:"PaymentConditionDescription,omitempty"`

	// PaymentReference:
	PaymentReference *string `json:"PaymentReference,omitempty"`

	// Remarks:
	Remarks *string `json:"Remarks,omitempty"`

	// SalesOrderLines:
	SalesOrderLines *[]byte `json:"SalesOrderLines,omitempty"`

	// Salesperson:
	Salesperson *types.GUID `json:"Salesperson,omitempty"`

	// SalespersonFullName:
	SalespersonFullName *string `json:"SalespersonFullName,omitempty"`

	// ShippingMethod:
	ShippingMethod *types.GUID `json:"ShippingMethod,omitempty"`

	// ShippingMethodDescription:
	ShippingMethodDescription *string `json:"ShippingMethodDescription,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// TaxSchedule:
	TaxSchedule *types.GUID `json:"TaxSchedule,omitempty"`

	// TaxScheduleCode:
	TaxScheduleCode *string `json:"TaxScheduleCode,omitempty"`

	// TaxScheduleDescription:
	TaxScheduleDescription *string `json:"TaxScheduleDescription,omitempty"`

	// WarehouseCode:
	WarehouseCode *string `json:"WarehouseCode,omitempty"`

	// WarehouseDescription:
	WarehouseDescription *string `json:"WarehouseDescription,omitempty"`

	// WarehouseID:
	WarehouseID *types.GUID `json:"WarehouseID,omitempty"`

	// YourRef:
	YourRef *string `json:"YourRef,omitempty"`
}

// List the SalesOrders entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *SalesOrdersEndpoint) List(ctx context.Context, division int, all bool) ([]*SalesOrders, error) {
	var entities []*SalesOrders
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/salesorder/SalesOrders?$select=*", division)
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
