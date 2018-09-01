// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package purchaseorder

import "github.com/mcnijman/go-exactonline/api"

type service struct {
	client *api.Client
}

// PurchaseOrderService is responsible for communication with the PurchaseOrder
// endpoints of the Exact Online API.
type PurchaseOrderService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	GoodsReceiptLines  *GoodsReceiptLinesEndpoint
	GoodsReceipts      *GoodsReceiptsEndpoint
	PurchaseOrderLines *PurchaseOrderLinesEndpoint
	PurchaseOrders     *PurchaseOrdersEndpoint
}

// NewPurchaseOrderService creates a new initialized instance of the
// PurchaseOrderService.
func NewPurchaseOrderService(apiClient *api.Client) *PurchaseOrderService {
	s := &PurchaseOrderService{client: apiClient}

	s.common.client = apiClient

	s.GoodsReceiptLines = (*GoodsReceiptLinesEndpoint)(&s.common)
	s.GoodsReceipts = (*GoodsReceiptsEndpoint)(&s.common)
	s.PurchaseOrderLines = (*PurchaseOrderLinesEndpoint)(&s.common)
	s.PurchaseOrders = (*PurchaseOrdersEndpoint)(&s.common)

	return s
}