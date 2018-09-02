// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package continuousmonitoring

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// IndicatorSignalsEndpoint is responsible for communicating with
// the IndicatorSignals endpoint of the ContinuousMonitoring service.
type IndicatorSignalsEndpoint service

// IndicatorSignals:
// Service: ContinuousMonitoring
// Entity: IndicatorSignals
// URL: /api/v1/beta/{division}/continuousmonitoring/IndicatorSignals
// HasWebhook: false
// IsInBeta: true
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ContinuousMonitoringIndicatorSignals
type IndicatorSignals struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Indicator signal description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EntryID: Primary key
	EntryID *types.GUID `json:"EntryID,omitempty"`

	// GLAccount: GL Account ID
	GLAccount *types.GUID `json:"GLAccount,omitempty"`

	// GlAccountCode: GL Account Code
	GlAccountCode *string `json:"GlAccountCode,omitempty"`

	// GlAccountDescription: GL Account Description
	GlAccountDescription *string `json:"GlAccountDescription,omitempty"`

	// Indicator: ID of Indicators
	Indicator *types.GUID `json:"Indicator,omitempty"`

	// IndicatorDescription: Indicator description
	IndicatorDescription *string `json:"IndicatorDescription,omitempty"`

	// IndicatorState: ID of IndicatorStates
	IndicatorState *types.GUID `json:"IndicatorState,omitempty"`

	// IndicatorStateReportingYear: Indicator states reporting year
	IndicatorStateReportingYear *int `json:"IndicatorStateReportingYear,omitempty"`

	// IndicatorType: Indicator type (1 = Balance G/L account per financial year, 2 = Usage of journals, 3 = Deviating amount entered, 4 = Liquidity, 5 = VAT Return deadline, 6 = Difference result in percentage, 7 = Different VAT code used)
	IndicatorType *int `json:"IndicatorType,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes of indicator signal
	Notes *string `json:"Notes,omitempty"`

	// NotesModified: Notes of indicator signal&#39;s modified date
	NotesModified *types.Date `json:"NotesModified,omitempty"`

	// NotesModifier: Notes of indicator signal&#39;s user ID modifier
	NotesModifier *types.GUID `json:"NotesModifier,omitempty"`

	// NotesModifierFullName: Name of modifier of notes in indicator signal
	NotesModifierFullName *string `json:"NotesModifierFullName,omitempty"`

	// Severity: Severity of the indicators (1 = Low, 2 = Medium, 3 = High, 4 = Critical)
	Severity *int `json:"Severity,omitempty"`

	// SignalDate: Signal&#39;s creation date
	SignalDate *types.Date `json:"SignalDate,omitempty"`

	// Status: Indicator signal status (0 = Ignore, 1 = Active, 2 = Archived)
	Status *int `json:"Status,omitempty"`
}

// List the IndicatorSignals entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *IndicatorSignalsEndpoint) List(ctx context.Context, division int, all bool) ([]*IndicatorSignals, error) {
	var entities []*IndicatorSignals
	u, err := s.client.ResolvePathWithDivision("/api/v1/beta/{division}/continuousmonitoring/IndicatorSignals?$select=*", division)
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
