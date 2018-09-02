// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package webhooks

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// WebhookSubscriptionsEndpoint is responsible for communicating with
// the WebhookSubscriptions endpoint of the Webhooks service.
type WebhookSubscriptionsEndpoint service

// WebhookSubscriptions:
// Service: Webhooks
// Entity: WebhookSubscriptions
// URL: /api/v1/{division}/webhooks/WebhookSubscriptions
// HasWebhook: false
// IsInBeta: false
// Methods: GET POST PUT DELETE
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=WebhooksWebhookSubscriptions
type WebhookSubscriptions struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// CallbackURL: Callback URL endpoint
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// ClientID: OAuth client Id
	ClientID *types.GUID `json:"ClientID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description of the OAuth Client
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Topic: Webhook subscription topic, e.g.: FinancialTransactions, Items, StockPositions
	Topic *string `json:"Topic,omitempty"`
}

// List the WebhookSubscriptions entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *WebhookSubscriptionsEndpoint) List(ctx context.Context, division int, all bool) ([]*WebhookSubscriptions, error) {
	var entities []*WebhookSubscriptions
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/webhooks/WebhookSubscriptions?$select=*", division)
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
