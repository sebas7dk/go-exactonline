// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package hrm

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// AbsenceRegistrationTransactionsEndpoint is responsible for communicating with
// the AbsenceRegistrationTransactions endpoint of the HRM service.
type AbsenceRegistrationTransactionsEndpoint service

// AbsenceRegistrationTransactions:
// Service: HRM
// Entity: AbsenceRegistrationTransactions
// URL: /api/v1/{division}/hrm/AbsenceRegistrationTransactions
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=HRMAbsenceRegistrationTransactions
type AbsenceRegistrationTransactions struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AbsenceRegistration: Reference key to Absence Registration
	AbsenceRegistration *types.GUID `json:"AbsenceRegistration,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndTime: End time on the last day of absence stored as DateTime, and the date should be ignored
	EndTime *types.Date `json:"EndTime,omitempty"`

	// ExpectedEndDate: Expected end date of absence
	ExpectedEndDate *types.Date `json:"ExpectedEndDate,omitempty"`

	// Hours: Total number of absence hours
	Hours *float64 `json:"Hours,omitempty"`

	// HoursFirstDay: Hours of absence on the first day
	HoursFirstDay *float64 `json:"HoursFirstDay,omitempty"`

	// HoursLastDay: Hours of absence on the last day
	HoursLastDay *float64 `json:"HoursLastDay,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Extra information for absence
	Notes *string `json:"Notes,omitempty"`

	// NotificationMoment: Notification moment of absence
	NotificationMoment *types.Date `json:"NotificationMoment,omitempty"`

	// PercentageDisablement: Percentage disablement
	PercentageDisablement *float64 `json:"PercentageDisablement,omitempty"`

	// StartDate: Start date of absence
	StartDate *types.Date `json:"StartDate,omitempty"`

	// StartTime: Start time on the first day of absence stored as DateTime, and the date should be ignored
	StartTime *types.Date `json:"StartTime,omitempty"`

	// Status: Status of absence, 0 = Open, 1 = Recovered
	Status *int `json:"Status,omitempty"`
}

// List the AbsenceRegistrationTransactions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AbsenceRegistrationTransactionsEndpoint) List(ctx context.Context, division int, all bool) ([]*AbsenceRegistrationTransactions, error) {
	var entities []*AbsenceRegistrationTransactions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/hrm/AbsenceRegistrationTransactions?$select=*", division)
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
