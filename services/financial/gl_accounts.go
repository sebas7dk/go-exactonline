// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// GLAccountsEndpoint is responsible for communicating with
// the GLAccounts endpoint of the Financial service.
type GLAccountsEndpoint service

// GLAccounts:
// Service: Financial
// Entity: GLAccounts
// URL: /api/v1/{division}/financial/GLAccounts
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialGLAccounts
type GLAccounts struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// AssimilatedVATBox:
	AssimilatedVATBox *int `json:"AssimilatedVATBox,omitempty"`

	// BalanceSide:
	BalanceSide *string `json:"BalanceSide,omitempty"`

	// BalanceType:
	BalanceType *string `json:"BalanceType,omitempty"`

	// BelcotaxType:
	BelcotaxType *int `json:"BelcotaxType,omitempty"`

	// Code:
	Code *string `json:"Code,omitempty"`

	// Compress:
	Compress *bool `json:"Compress,omitempty"`

	// Costcenter:
	Costcenter *string `json:"Costcenter,omitempty"`

	// CostcenterDescription:
	CostcenterDescription *string `json:"CostcenterDescription,omitempty"`

	// Costunit:
	Costunit *string `json:"Costunit,omitempty"`

	// CostunitDescription:
	CostunitDescription *string `json:"CostunitDescription,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// ExcludeVATListing:
	ExcludeVATListing *byte `json:"ExcludeVATListing,omitempty"`

	// ExpenseNonDeductiblePercentage:
	ExpenseNonDeductiblePercentage *float64 `json:"ExpenseNonDeductiblePercentage,omitempty"`

	// IsBlocked:
	IsBlocked *bool `json:"IsBlocked,omitempty"`

	// Matching:
	Matching *bool `json:"Matching,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// PrivateGLAccount:
	PrivateGLAccount *types.GUID `json:"PrivateGLAccount,omitempty"`

	// PrivatePercentage:
	PrivatePercentage *float64 `json:"PrivatePercentage,omitempty"`

	// ReportingCode:
	ReportingCode *string `json:"ReportingCode,omitempty"`

	// RevalueCurrency:
	RevalueCurrency *bool `json:"RevalueCurrency,omitempty"`

	// SearchCode:
	SearchCode *string `json:"SearchCode,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// UseCostcenter:
	UseCostcenter *byte `json:"UseCostcenter,omitempty"`

	// UseCostunit:
	UseCostunit *byte `json:"UseCostunit,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATDescription:
	VATDescription *string `json:"VATDescription,omitempty"`

	// VATGLAccountType:
	VATGLAccountType *string `json:"VATGLAccountType,omitempty"`

	// VATNonDeductibleGLAccount:
	VATNonDeductibleGLAccount *types.GUID `json:"VATNonDeductibleGLAccount,omitempty"`

	// VATNonDeductiblePercentage:
	VATNonDeductiblePercentage *float64 `json:"VATNonDeductiblePercentage,omitempty"`

	// VATSystem:
	VATSystem *string `json:"VATSystem,omitempty"`

	// YearEndCostGLAccount:
	YearEndCostGLAccount *types.GUID `json:"YearEndCostGLAccount,omitempty"`

	// YearEndReflectionGLAccount:
	YearEndReflectionGLAccount *types.GUID `json:"YearEndReflectionGLAccount,omitempty"`
}

// List the GLAccounts entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *GLAccountsEndpoint) List(ctx context.Context, division int, all bool) ([]*GLAccounts, error) {
	var entities []*GLAccounts
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financial/GLAccounts?$select=*", division)
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
