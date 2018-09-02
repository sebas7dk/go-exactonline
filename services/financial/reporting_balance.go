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

// ReportingBalanceEndpoint is responsible for communicating with
// the ReportingBalance endpoint of the Financial service.
type ReportingBalanceEndpoint service

// ReportingBalance:
// Service: Financial
// Entity: ReportingBalance
// URL: /api/v1/{division}/financial/ReportingBalance
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=FinancialReportingBalance
type ReportingBalance struct {
	// ID: Record ID
	ID *int64 `json:"ID,omitempty"`

	// Amount: The sum of the amounts of all transactions in the grouping.
	Amount *float64 `json:"Amount,omitempty"`

	// AmountCredit: The sum of the amounts of all credit transactions in the grouping.
	AmountCredit *float64 `json:"AmountCredit,omitempty"`

	// AmountDebit: The sum of the amounts of all debit transactions in the grouping.
	AmountDebit *float64 `json:"AmountDebit,omitempty"`

	// BalanceType: Balance type of the G/L account: B = Balance Sheet, W = Profit &amp; Loss.
	BalanceType *string `json:"BalanceType,omitempty"`

	// CostCenterCode: The code of the cost center.
	CostCenterCode *string `json:"CostCenterCode,omitempty"`

	// CostCenterDescription: The description of the cost center.
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostUnitCode: The code of the cost unit.
	CostUnitCode *string `json:"CostUnitCode,omitempty"`

	// CostUnitDescription: The description of the cost unit.
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// Count: The number of transactions in the grouping.
	Count *int `json:"Count,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// GLAccount: G/L account
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountCode: The code of the G/L account.
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription: The description of the G/L account.
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// ReportingPeriod: The reporting period of the transactions in the grouping.
	ReportingPeriod *int `json:"ReportingPeriod,omitempty"`

	// ReportingYear: The reporting year of the transactions in the grouping.
	ReportingYear *int `json:"ReportingYear,omitempty"`

	// Status: Status: 20 = Open, 50 = Processed. To get &#39;after entry&#39; results, both Open and Processed amounts have to be included. This is by default, so it requires no extra filtering.
	Status *int `json:"Status,omitempty"`

	// Type: The type of the transactions in the grouping.
	Type *int `json:"Type,omitempty"`
}

// List the ReportingBalance entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *ReportingBalanceEndpoint) List(ctx context.Context, division int, all bool) ([]*ReportingBalance, error) {
	var entities []*ReportingBalance
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/financial/ReportingBalance?$select=*", division)
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
