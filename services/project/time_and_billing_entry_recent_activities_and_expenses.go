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

// TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint is responsible for communicating with
// the TimeAndBillingEntryRecentActivitiesAndExpenses endpoint of the Project service.
type TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint service

// TimeAndBillingEntryRecentActivitiesAndExpenses:
// Service: Project
// Entity: TimeAndBillingEntryRecentActivitiesAndExpenses
// URL: /api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadProjectTimeAndBillingEntryRecentActivitiesAndExpenses
type TimeAndBillingEntryRecentActivitiesAndExpenses struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// DateLastUsed: Date last used
	DateLastUsed *types.Date `json:"DateLastUsed,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// ParentDescription: Description of Parent
	ParentDescription *string `json:"ParentDescription,omitempty"`
}

// List the TimeAndBillingEntryRecentActivitiesAndExpenses entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *TimeAndBillingEntryRecentActivitiesAndExpensesEndpoint) List(ctx context.Context, division int, all bool) ([]*TimeAndBillingEntryRecentActivitiesAndExpenses, error) {
	var entities []*TimeAndBillingEntryRecentActivitiesAndExpenses
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/project/TimeAndBillingEntryRecentActivitiesAndExpenses?$select=*", division)
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
