// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package crm

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// OpportunitiesEndpoint is responsible for communicating with
// the Opportunities endpoint of the CRM service.
type OpportunitiesEndpoint service

// Opportunities:
// Service: CRM
// Entity: Opportunities
// URL: /api/v1/{division}/crm/Opportunities
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=CRMOpportunities
type Opportunities struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Account:
	Account *types.GUID `json:"Account,omitempty"`

	// Accountant:
	Accountant *types.GUID `json:"Accountant,omitempty"`

	// AccountantCode:
	AccountantCode *string `json:"AccountantCode,omitempty"`

	// AccountantName:
	AccountantName *string `json:"AccountantName,omitempty"`

	// AccountCode:
	AccountCode *string `json:"AccountCode,omitempty"`

	// AccountName:
	AccountName *string `json:"AccountName,omitempty"`

	// ActionDate:
	ActionDate *types.Date `json:"ActionDate,omitempty"`

	// AmountDC:
	AmountDC *float64 `json:"AmountDC,omitempty"`

	// AmountFC:
	AmountFC *float64 `json:"AmountFC,omitempty"`

	// Campaign:
	Campaign *types.GUID `json:"Campaign,omitempty"`

	// CampaignDescription:
	CampaignDescription *string `json:"CampaignDescription,omitempty"`

	// Channel:
	Channel *int `json:"Channel,omitempty"`

	// ChannelDescription:
	ChannelDescription *string `json:"ChannelDescription,omitempty"`

	// CloseDate:
	CloseDate *types.Date `json:"CloseDate,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Currency:
	Currency *string `json:"Currency,omitempty"`

	// Division:
	Division *int `json:"Division,omitempty"`

	// LeadSource:
	LeadSource *types.GUID `json:"LeadSource,omitempty"`

	// LeadSourceDescription:
	LeadSourceDescription *string `json:"LeadSourceDescription,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Name:
	Name *string `json:"Name,omitempty"`

	// NextAction:
	NextAction *string `json:"NextAction,omitempty"`

	// Notes:
	Notes *string `json:"Notes,omitempty"`

	// OpportunityDepartmentCode:
	OpportunityDepartmentCode *int `json:"OpportunityDepartmentCode,omitempty"`

	// OpportunityDepartmentDescription:
	OpportunityDepartmentDescription *string `json:"OpportunityDepartmentDescription,omitempty"`

	// OpportunityStage:
	OpportunityStage *types.GUID `json:"OpportunityStage,omitempty"`

	// OpportunityStageDescription:
	OpportunityStageDescription *string `json:"OpportunityStageDescription,omitempty"`

	// OpportunityStatus:
	OpportunityStatus *int `json:"OpportunityStatus,omitempty"`

	// OpportunityType:
	OpportunityType *int `json:"OpportunityType,omitempty"`

	// OpportunityTypeDescription:
	OpportunityTypeDescription *string `json:"OpportunityTypeDescription,omitempty"`

	// Owner:
	Owner *types.GUID `json:"Owner,omitempty"`

	// OwnerFullName:
	OwnerFullName *string `json:"OwnerFullName,omitempty"`

	// Probability:
	Probability *float64 `json:"Probability,omitempty"`

	// Project:
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectCode:
	ProjectCode *string `json:"ProjectCode,omitempty"`

	// ProjectDescription:
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// RateFC:
	RateFC *float64 `json:"RateFC,omitempty"`

	// ReasonCode:
	ReasonCode *types.GUID `json:"ReasonCode,omitempty"`

	// ReasonCodeDescription:
	ReasonCodeDescription *string `json:"ReasonCodeDescription,omitempty"`

	// Reseller:
	Reseller *types.GUID `json:"Reseller,omitempty"`

	// ResellerCode:
	ResellerCode *string `json:"ResellerCode,omitempty"`

	// ResellerName:
	ResellerName *string `json:"ResellerName,omitempty"`

	// SalesType:
	SalesType *types.GUID `json:"SalesType,omitempty"`

	// SalesTypeDescription:
	SalesTypeDescription *string `json:"SalesTypeDescription,omitempty"`
}

// List the Opportunities entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *OpportunitiesEndpoint) List(ctx context.Context, division int, all bool) ([]*Opportunities, error) {
	var entities []*Opportunities
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/crm/Opportunities?$select=*", division)
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
