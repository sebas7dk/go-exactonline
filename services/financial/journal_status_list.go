// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import "context"

// JournalStatusListEndpoint is responsible for communicating with
// the JournalStatusList endpoint of the Financial service.
type JournalStatusListEndpoint service

// JournalStatusList:
// Service: Financial
// Entity: JournalStatusList
// URL: /api/v1/{division}/read/financial/JournalStatusList
// HasWebhook: true
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ReadFinancialJournalStatusList
type JournalStatusList struct {
	// Journal:
	Journal *string `json:"Journal,omitempty"`

	// Period:
	Period *int `json:"Period,omitempty"`

	// Year:
	Year *int `json:"Year,omitempty"`

	// JournalDescription:
	JournalDescription *string `json:"JournalDescription,omitempty"`

	// JournalType:
	JournalType *int `json:"JournalType,omitempty"`

	// JournalTypeDescription:
	JournalTypeDescription *string `json:"JournalTypeDescription,omitempty"`

	// Status:
	Status *int `json:"Status,omitempty"`

	// StatusDescription:
	StatusDescription *string `json:"StatusDescription,omitempty"`
}

func (s *JournalStatusList) GetIdentifier() string {
	return *s.Journal
}

// List the JournalStatusList entities in the provided divison.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *JournalStatusListEndpoint) List(ctx context.Context, division int, all bool) ([]*JournalStatusList, error) {
	var entities []*JournalStatusList
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/JournalStatusList?$select=*", division)
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