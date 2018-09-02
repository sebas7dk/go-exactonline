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

// InvoiceTermsEndpoint is responsible for communicating with
// the InvoiceTerms endpoint of the Project service.
type InvoiceTermsEndpoint service

// InvoiceTerms:
// Service: Project
// Entity: InvoiceTerms
// URL: /api/v1/{division}/project/InvoiceTerms
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ProjectInvoiceTerms
type InvoiceTerms struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Amount: Amount in the currency of the transaction
	Amount *float64 `json:"Amount,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Deliverable: WBS&#39;s deliverable linked to invoice term
	Deliverable *string `json:"Deliverable,omitempty"`

	// Description: Description of invoice term
	Description *string `json:"Description,omitempty"`

	// Division: Division number
	Division *int `json:"Division,omitempty"`

	// ExecutionFromDate: Execution date: From
	ExecutionFromDate *types.Date `json:"ExecutionFromDate,omitempty"`

	// ExecutionToDate: Execution date: To
	ExecutionToDate *types.Date `json:"ExecutionToDate,omitempty"`

	// InvoiceDate: Invoice date
	InvoiceDate *types.Date `json:"InvoiceDate,omitempty"`

	// Item: Reference to item
	Item *types.GUID `json:"Item,omitempty"`

	// ItemDescription: Description of item
	ItemDescription *string `json:"ItemDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`

	// Percentage: Percentage of amount per project&#39;s budgeted amount
	Percentage *float64 `json:"Percentage,omitempty"`

	// Project: Reference to project
	Project *types.GUID `json:"Project,omitempty"`

	// ProjectDescription: Description of project
	ProjectDescription *string `json:"ProjectDescription,omitempty"`

	// VATCode: Reference to VATCode
	VATCode *string `json:"VATCode,omitempty"`

	// VATCodeDescription: Description of VATCode
	VATCodeDescription *string `json:"VATCodeDescription,omitempty"`

	// VATPercentage: VATCode percentage
	VATPercentage *float64 `json:"VATPercentage,omitempty"`
}

// List the InvoiceTerms entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *InvoiceTermsEndpoint) List(ctx context.Context, division int, all bool) ([]*InvoiceTerms, error) {
	var entities []*InvoiceTerms
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/project/InvoiceTerms?$select=*", division)
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
