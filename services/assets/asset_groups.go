// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package assets

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// AssetGroupsEndpoint is responsible for communicating with
// the AssetGroups endpoint of the Assets service.
type AssetGroupsEndpoint service

// AssetGroups:
// Service: Assets
// Entity: AssetGroups
// URL: /api/v1/{division}/assets/AssetGroups
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AssetsAssetGroups
type AssetGroups struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// Code: Code of the asset group
	Code *string `json:"Code,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// DepreciationMethod: Default depreciation method of the assets in this asset group
	DepreciationMethod *types.GUID `json:"DepreciationMethod,omitempty"`

	// DepreciationMethodCode: Code of the depreciation method
	DepreciationMethodCode *string `json:"DepreciationMethodCode,omitempty"`

	// DepreciationMethodDescription: Description of the depreciation method
	DepreciationMethodDescription *string `json:"DepreciationMethodDescription,omitempty"`

	// Description: Description of the asset group
	Description *string `json:"Description,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// GLAccountAssets: GLAccount for the assets
	GLAccountAssets *types.GUID `json:"GLAccountAssets,omitempty"`

	// GLAccountAssetsCode: Code of the GLAccount for the assets
	GLAccountAssetsCode *string `json:"GLAccountAssetsCode,omitempty"`

	// GLAccountAssetsDescription: Description of the GLAccount for the assets
	GLAccountAssetsDescription *string `json:"GLAccountAssetsDescription,omitempty"`

	// GLAccountDepreciationBS: GLAccount for depreciation (Balance sheet)
	GLAccountDepreciationBS *types.GUID `json:"GLAccountDepreciationBS,omitempty"`

	// GLAccountDepreciationBSCode: Code of the GLAccount for depreciation (Balance sheet)
	GLAccountDepreciationBSCode *string `json:"GLAccountDepreciationBSCode,omitempty"`

	// GLAccountDepreciationBSDescription: Description of the GLAccount for depreciation (Balance sheet)
	GLAccountDepreciationBSDescription *string `json:"GLAccountDepreciationBSDescription,omitempty"`

	// GLAccountDepreciationPL: GLAccount for depreciation (Profit &amp; Loss)
	GLAccountDepreciationPL *types.GUID `json:"GLAccountDepreciationPL,omitempty"`

	// GLAccountDepreciationPLCode: Code of the GLAccount for depreciation (Profit &amp; Loss)
	GLAccountDepreciationPLCode *string `json:"GLAccountDepreciationPLCode,omitempty"`

	// GLAccountDepreciationPLDescription: Description of the GLAccount for depreciation (Profit &amp; Loss)
	GLAccountDepreciationPLDescription *string `json:"GLAccountDepreciationPLDescription,omitempty"`

	// GLAccountRevaluationBS: GLAccount for revaluation (Balance sheet)
	GLAccountRevaluationBS *types.GUID `json:"GLAccountRevaluationBS,omitempty"`

	// GLAccountRevaluationBSCode: Code of the GLAccount for revaluation (Balance sheet)
	GLAccountRevaluationBSCode *string `json:"GLAccountRevaluationBSCode,omitempty"`

	// GLAccountRevaluationBSDescription: Description of the GLAccount for revaluation (Balance sheet)
	GLAccountRevaluationBSDescription *string `json:"GLAccountRevaluationBSDescription,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// Notes: Notes
	Notes *string `json:"Notes,omitempty"`
}

// List the AssetGroups entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *AssetGroupsEndpoint) List(ctx context.Context, division int, all bool) ([]*AssetGroups, error) {
	var entities []*AssetGroups
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/assets/AssetGroups?$select=*", division)
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
