// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package project

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// RecentHoursEndpoint is responsible for communicating with
// the RecentHours endpoint of the Project service.
type RecentHoursEndpoint service

// RecentHours:
// Service: Project
// Entity: RecentHours
// URL: /api/v1/{division}/read/project/RecentHours
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectRecentHours
type RecentHours struct {
	// Id: Primary key
	Id *int `json:"Id,omitempty"`

	// AccountCode: Code of Account
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountId: Reference to Account
	AccountId *types.GUID `json:"AccountId,omitempty"`

	// AccountName: Name of Account
	AccountName *string `json:"AccountName,omitempty"`

	// Activity: Reference to Activity
	Activity *types.GUID `json:"Activity,omitempty"`

	// ActivityDescription: Description of Activity
	ActivityDescription *string `json:"ActivityDescription,omitempty"`

	// Date: Date
	Date *types.Date `json:"Date,omitempty"`

	// EntryId: Entry ID
	EntryId *types.GUID `json:"EntryId,omitempty"`

	// HoursApproved: Hours approved
	HoursApproved *float64 `json:"HoursApproved,omitempty"`

	// HoursApprovedBillable: Billable hours approved
	HoursApprovedBillable *float64 `json:"HoursApprovedBillable,omitempty"`

	// HoursDraft: Hours draft
	HoursDraft *float64 `json:"HoursDraft,omitempty"`

	// HoursDraftBillable: Billable hours draft
	HoursDraftBillable *float64 `json:"HoursDraftBillable,omitempty"`

	// HoursRejected: Hours rejected
	HoursRejected *float64 `json:"HoursRejected,omitempty"`

	// HoursRejectedBillable: Billable hours rejected
	HoursRejectedBillable *float64 `json:"HoursRejectedBillable,omitempty"`

	// HoursSubmitted: Hours submitted
	HoursSubmitted *float64 `json:"HoursSubmitted,omitempty"`

	// HoursSubmittedBillable: Billable hours submitted
	HoursSubmittedBillable *float64 `json:"HoursSubmittedBillable,omitempty"`

	// ItemCode: Code of Item
	ItemCode *string `json:"ItemCode,omitempty"`

	// ItemDescription: Description of Item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// ItemId: Reference to Item
	ItemId *types.GUID `json:"ItemId,omitempty"`

	// Notes: Remarks
	Notes *string `json:"Notes,omitempty"`

	// ProjectCode: Code of Project
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription: Description of Project
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// ProjectId: Reference to Project
	ProjectId *types.GUID `json:"ProjectId,omitempty"`

	// WeekNumber: Week number
	WeekNumber *int `json:"WeekNumber,omitempty"`
}

func (s *RecentHours) GetIdentifier() int {
	return *s.Id
}

// List the RecentHours entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *RecentHoursEndpoint) List(ctx context.Context, division int, all bool) ([]*RecentHours, error) {
	var entities []*RecentHours
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/RecentHours?$select=*", division)
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