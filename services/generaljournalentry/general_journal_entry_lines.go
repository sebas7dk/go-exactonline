// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package generaljournalentry

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// GeneralJournalEntryLinesEndpoint is responsible for communicating with
// the GeneralJournalEntryLines endpoint of the GeneralJournalEntry service.
type GeneralJournalEntryLinesEndpoint service

// GeneralJournalEntryLines:
// Service: GeneralJournalEntry
// Entity: GeneralJournalEntryLines
// URL: /api/v1/{division}/generaljournalentry/GeneralJournalEntryLines
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=GeneralJournalEntryGeneralJournalEntryLines
type GeneralJournalEntryLines struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// AmountVATDC:
	AmountVATDC *float64 `json:"AmountVATDC,omitempty"`

	// AmountVATFC:
	AmountVATFC *float64 `json:"AmountVATFC,omitempty"`

	// Asset:
	Asset *types.GUID `json:"Asset,omitempty"`

	// AssetCode:
	AssetCode *string `json:"AssetCode,omitempty"`

	// AssetDescription:
	AssetDescription *string `json:"AssetDescription,omitempty"`

	// CostCenter:
	CostCenter *string `json:"CostCenter,omitempty"`

	// CostCenterDescription:
	CostCenterDescription *string `json:"CostCenterDescription,omitempty"`

	// CostUnit:
	CostUnit *string `json:"CostUnit,omitempty"`

	// CostUnitDescription:
	CostUnitDescription *string `json:"CostUnitDescription,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Date:
	Date *types.Date `json:"Date,omitempty"`

	// Description:
	Description *string `json:"Description,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// Document:
	Document *types.GUID `json:"Document,omitempty"`

	// DocumentNumber:
	DocumentNumber *int `json:"DocumentNumber,omitempty"`

	// DocumentSubject:
	DocumentSubject *string `json:"DocumentSubject,omitempty"`

	// EntryID:
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// EntryNumber:
	EntryNumber *int `json:"EntryNumber,omitempty"`

	// GLAccount:
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GLAccountCode:
	GLAccountCode *string `json:"GLAccountCode,omitempty"`

	// GLAccountDescription:
	GLAccountDescription *string `json:"GLAccountDescription,omitempty"`

	// LineNumber:
	LineNumber *int `json:"LineNumber,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OffsetID:
	OffsetID *types.GUID `json:"OffsetID,omitempty"`

	// OurRef:
	OurRef *int `json:"OurRef,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// VATBaseAmountDC:
	VATBaseAmountDC *float64 `json:"VATBaseAmountDC,omitempty"`

	// VATBaseAmountFC:
	VATBaseAmountFC *float64 `json:"VATBaseAmountFC,omitempty"`

	// VATCode:
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription:
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage:
	VATPercentage *float64 `json:"VATPercentage,omitempty"`

	// VATType:
	VATType *string `json:"VATType,omitempty"`
}

// List the GeneralJournalEntryLines entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *GeneralJournalEntryLinesEndpoint) List(ctx context.Context, division int, all bool) ([]*GeneralJournalEntryLines, error) {
	var entities []*GeneralJournalEntryLines
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/generaljournalentry/GeneralJournalEntryLines?$select=*", division)
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
