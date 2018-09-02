// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package mailbox

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// MailMessagesReceivedEndpoint is responsible for communicating with
// the MailMessagesReceived endpoint of the Mailbox service.
type MailMessagesReceivedEndpoint service

// MailMessagesReceived:
// Service: Mailbox
// Entity: MailMessagesReceived
// URL: /api/v1/{division}/mailbox/MailMessagesReceived
// HasWebhook: true
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=MailboxMailMessagesReceived
type MailMessagesReceived struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Bank:
	Bank *types.GUID `json:"Bank,omitempty"`

	// BankAccount:
	BankAccount *string `json:"BankAccount,omitempty"`

	// Created:
	Created *types.Date `json:"Created,omitempty"`

	// Creator:
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName:
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// ForDivision:
	ForDivision *int `json:"ForDivision,omitempty"`

	// Modified:
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier:
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName:
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Operation:
	Operation *int `json:"Operation,omitempty"`

	// OriginalMessage:
	OriginalMessage *types.GUID `json:"OriginalMessage,omitempty"`

	// OriginalMessageSubject:
	OriginalMessageSubject *string `json:"OriginalMessageSubject,omitempty"`

	// PartnerKey:
	PartnerKey *types.GUID `json:"PartnerKey,omitempty"`

	// Quantity:
	Quantity *float64 `json:"Quantity,omitempty"`

	// RecipientAccount:
	RecipientAccount *types.GUID `json:"RecipientAccount,omitempty"`

	// RecipientDeleted:
	RecipientDeleted *byte `json:"RecipientDeleted,omitempty"`

	// RecipientMailbox:
	RecipientMailbox *string `json:"RecipientMailbox,omitempty"`

	// RecipientMailboxDescription:
	RecipientMailboxDescription *string `json:"RecipientMailboxDescription,omitempty"`

	// RecipientMailboxID:
	RecipientMailboxID *types.GUID `json:"RecipientMailboxID,omitempty"`

	// RecipientStatus:
	RecipientStatus *int `json:"RecipientStatus,omitempty"`

	// RecipientStatusDescription:
	RecipientStatusDescription *string `json:"RecipientStatusDescription,omitempty"`

	// SenderAccount:
	SenderAccount *types.GUID `json:"SenderAccount,omitempty"`

	// SenderDateSent:
	SenderDateSent *types.Date `json:"SenderDateSent,omitempty"`

	// SenderDeleted:
	SenderDeleted *byte `json:"SenderDeleted,omitempty"`

	// SenderIPAddress:
	SenderIPAddress *string `json:"SenderIPAddress,omitempty"`

	// SenderMailbox:
	SenderMailbox *string `json:"SenderMailbox,omitempty"`

	// SenderMailboxDescription:
	SenderMailboxDescription *string `json:"SenderMailboxDescription,omitempty"`

	// SenderMailboxID:
	SenderMailboxID *types.GUID `json:"SenderMailboxID,omitempty"`

	// Subject:
	Subject *string `json:"Subject,omitempty"`

	// SynchronizationCode:
	SynchronizationCode *string `json:"SynchronizationCode,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`
}

// List the MailMessagesReceived entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MailMessagesReceivedEndpoint) List(ctx context.Context, division int, all bool) ([]*MailMessagesReceived, error) {
	var entities []*MailMessagesReceived
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/MailMessagesReceived?$select=*", division)
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
