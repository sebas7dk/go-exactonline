// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package users

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// UserRolesPerDivisionEndpoint is responsible for communicating with
// the UserRolesPerDivision endpoint of the Users service.
type UserRolesPerDivisionEndpoint service

// UserRolesPerDivision:
// Service: Users
// Entity: UserRolesPerDivision
// URL: /api/v1/{division}/users/UserRolesPerDivision
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=UsersUserRolesPerDivision
type UserRolesPerDivision struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of the creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of the creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Description: Description
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// EndDate: Indicates the date and time when te role becomes inactive for the user
	EndDate *types.Date `json:"EndDate,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of the last modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of the last modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Role: The role that the user is linked to
	Role *int `json:"Role,omitempty"`

	// RoleLevel: Rolelevel sets the level on which a role for a user is active. This can be: 1 = Database, 2 = Customer, 3 = Division, 100 = Transferred to accountant
	RoleLevel *int `json:"RoleLevel,omitempty"`

	// StartDate: Indicates the date when the role becomes active for the user
	StartDate *types.Date `json:"StartDate,omitempty"`

	// UserID: The user that is linked to the role
	UserID *types.GUID `json:"UserID,omitempty"`
}

// List the UserRolesPerDivision entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *UserRolesPerDivisionEndpoint) List(ctx context.Context, division int, all bool) ([]*UserRolesPerDivision, error) {
	var entities []*UserRolesPerDivision
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/users/UserRolesPerDivision?$select=*", division)
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
