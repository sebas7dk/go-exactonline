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

// MailMessageAttachmentsEndpoint is responsible for communicating with
// the MailMessageAttachments endpoint of the Mailbox service.
type MailMessageAttachmentsEndpoint service

// MailMessageAttachments:
// Service: Mailbox
// Entity: MailMessageAttachments
// URL: /api/v1/{division}/mailbox/MailMessageAttachments
// HasWebhook: true
// IsInBeta: false
// Methods: GET POST
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=MailboxMailMessageAttachments
type MailMessageAttachments struct {
	// ID:
	ID *types.GUID `json:"ID,omitempty"`

	// Attachment:
	Attachment *[]byte `json:"Attachment,omitempty"`

	// AttachmentFileExtension:
	AttachmentFileExtension *string `json:"AttachmentFileExtension,omitempty"`

	// AttachmentFileName:
	AttachmentFileName *string `json:"AttachmentFileName,omitempty"`

	// FileSize:
	FileSize *int64 `json:"FileSize,omitempty"`

	// MailMessageID:
	MailMessageID *types.GUID `json:"MailMessageID,omitempty"`

	// RecipientAccount:
	RecipientAccount *types.GUID `json:"RecipientAccount,omitempty"`

	// SenderAccount:
	SenderAccount *types.GUID `json:"SenderAccount,omitempty"`

	// Type:
	Type *int `json:"Type,omitempty"`

	// TypeDescription:
	TypeDescription *string `json:"TypeDescription,omitempty"`

	// Url:
	Url *string `json:"Url,omitempty"`
}

// List the MailMessageAttachments entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *MailMessageAttachmentsEndpoint) List(ctx context.Context, division int, all bool) ([]*MailMessageAttachments, error) {
	var entities []*MailMessageAttachments
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/mailbox/MailMessageAttachments?$select=*", division)
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
