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

// ShopOrderRoutingStepPlansEndpoint is responsible for communicating with
// the ShopOrderRoutingStepPlans endpoint of the Manufacturing service.
type ShopOrderRoutingStepPlansEndpoint service

// ShopOrderRoutingStepPlans:
// Service: Manufacturing
// Entity: ShopOrderRoutingStepPlans
// URL: /api/v1/{division}/manufacturing/ShopOrderRoutingStepPlans
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ManufacturingShopOrderRoutingStepPlans
type ShopOrderRoutingStepPlans struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Account: Reference to Account providing the Outsourced item
	Account *types.GUID `json:"Account,omitempty"`

	// AccountName: Account name
	AccountName *string `json:"AccountName,omitempty"`

	// AccountNumber: Account number
	AccountNumber *string `json:"AccountNumber,omitempty"`

	// AttendedPercentage: Attended Percentage
	AttendedPercentage *float64 `json:"AttendedPercentage,omitempty"`

	// Backflush: Indicates if this is a backflush step
	Backflush *byte `json:"Backflush,omitempty"`

	// CostPerItem: Total cost / Shop order planned quantity
	CostPerItem *float64 `json:"CostPerItem,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the operation
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EfficiencyPercentage: Efficiency Percentage
	EfficiencyPercentage *float64 `json:"EfficiencyPercentage,omitempty"`

	// FactorType: Conversion factor type between Shop order Item and Subcontract purchase Unit
	FactorType *int `json:"FactorType,omitempty"`

	// LineNumber: Sequential order of the operation
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// Operation: Reference to Operations
	Operation *types.GUID `json:"Operation,omitempty"`

	// OperationCode: Code of the routing step operation
	OperationCode *string `json:"OperationCode,omitempty"`

	// OperationDescription: Description of the operation step
	OperationDescription *string `json:"OperationDescription,omitempty"`

	// OperationResource: Reference to OperationResources
	OperationResource *types.GUID `json:"OperationResource,omitempty"`

	// PlannedEndDate: Planned end date
	PlannedEndDate *types.Date `json:"PlannedEndDate,omitempty"`

	// PlannedRunHours: Planned run hours
	PlannedRunHours *float64 `json:"PlannedRunHours,omitempty"`

	// PlannedSetupHours: Planned setup hours
	PlannedSetupHours *float64 `json:"PlannedSetupHours,omitempty"`

	// PlannedStartDate: Planned start date
	PlannedStartDate *types.Date `json:"PlannedStartDate,omitempty"`

	// PlannedTotalHours: Setup hours &#43; Run hours
	PlannedTotalHours *float64 `json:"PlannedTotalHours,omitempty"`

	// PurchaseUnit: Reference to Units
	PurchaseUnit *string `json:"PurchaseUnit,omitempty"`

	// PurchaseUnitFactor: Purchase Unit Factor
	PurchaseUnitFactor *float64 `json:"PurchaseUnitFactor,omitempty"`

	// PurchaseUnitPriceFC: Purchase Unit Price in the currency of the transaction
	PurchaseUnitPriceFC *float64 `json:"PurchaseUnitPriceFC,omitempty"`

	// PurchaseUnitQuantity: Purchase unit quantity of the plan
	PurchaseUnitQuantity *float64 `json:"PurchaseUnitQuantity,omitempty"`

	// RoutingStepType: Reference to RoutingStepTypes
	RoutingStepType *int `json:"RoutingStepType,omitempty"`

	// Run: Used in conjunction with RunMethod, and EfficiencyPercentage to determine PlannedRunHours
	Run *float64 `json:"Run,omitempty"`

	// RunMethod: Reference to OperationMethod
	RunMethod *int `json:"RunMethod,omitempty"`

	// RunMethodDescription: Description of RunMethod
	RunMethodDescription *string `json:"RunMethodDescription,omitempty"`

	// Setup: Used in conjunction with SetupCount and Setup Unit to determine PlannedSetupHours
	Setup *float64 `json:"Setup,omitempty"`

	// SetupUnit: Reference to TimeUnits
	SetupUnit *string `json:"SetupUnit,omitempty"`

	// ShopOrder: Reference to Shop orders
	ShopOrder *types.GUID `json:"ShopOrder,omitempty"`

	// Status: Reference to OperationStatus
	Status *int `json:"Status,omitempty"`

	// StatusDescription: Description of Status
	StatusDescription *string `json:"StatusDescription,omitempty"`

	// SubcontractedLeadDays: Subcontracted lead days
	SubcontractedLeadDays *int `json:"SubcontractedLeadDays,omitempty"`

	// TimeTransactions: Collection of TimeTransactions
	TimeTransactions *[]byte `json:"TimeTransactions,omitempty"`

	// TotalCostDC: Total cost of the routing line
	TotalCostDC *float64 `json:"TotalCostDC,omitempty"`

	// Workcenter: Reference to Workcenters
	Workcenter *types.GUID `json:"Workcenter,omitempty"`

	// WorkcenterCode: Workcenter code
	WorkcenterCode *string `json:"WorkcenterCode,omitempty"`

	// WorkcenterDescription: Workcenter description
	WorkcenterDescription *string `json:"WorkcenterDescription,omitempty"`
}

// List the ShopOrderRoutingStepPlans entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ShopOrderRoutingStepPlansEndpoint) List(ctx context.Context, division int, all bool) ([]*ShopOrderRoutingStepPlans, error) {
	var entities []*ShopOrderRoutingStepPlans
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/manufacturing/ShopOrderRoutingStepPlans?$select=*", division)
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
