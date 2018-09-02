// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package payroll

import (
	"context"

	"github.com/mcnijman/go-exactonline/types"
)

// EmploymentSalariesEndpoint is responsible for communicating with
// the EmploymentSalaries endpoint of the Payroll service.
type EmploymentSalariesEndpoint service

// EmploymentSalaries:
// Service: Payroll
// Entity: EmploymentSalaries
// URL: /api/v1/{division}/payroll/EmploymentSalaries
// HasWebhook: false
// IsInBeta: false
// Methods: GET
// Endpoint docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=PayrollEmploymentSalaries
type EmploymentSalaries struct {
	// ID: Primary key
	ID *types.GUID `json:"ID,omitempty"`

	// AverageDaysPerWeek: The average number of contract days that an employee works per week
	AverageDaysPerWeek *float64 `json:"AverageDaysPerWeek,omitempty"`

	// AverageHoursPerWeek: The average number of contract hours that an employee works per week
	AverageHoursPerWeek *float64 `json:"AverageHoursPerWeek,omitempty"`

	// Created: Creation date
	Created *types.Date `json:"Created,omitempty"`

	// Creator: User ID of creator
	Creator *types.GUID `json:"Creator,omitempty"`

	// CreatorFullName: Name of creator
	CreatorFullName *string `json:"CreatorFullName,omitempty"`

	// Division: Division code
	Division *int `json:"Division,omitempty"`

	// Employee: Employee ID
	Employee *types.GUID `json:"Employee,omitempty"`

	// EmployeeFullName: Name of employee
	EmployeeFullName *string `json:"EmployeeFullName,omitempty"`

	// EmployeeHID: Employee number
	EmployeeHID *int `json:"EmployeeHID,omitempty"`

	// Employment: Employment
	Employment *types.GUID `json:"Employment,omitempty"`

	// EmploymentHID: Employment number
	EmploymentHID *int `json:"EmploymentHID,omitempty"`

	// EmploymentSalaryType: Salary type of employment. 1 - Periodical (fixed), 2 - Per hour (variable)
	EmploymentSalaryType *int `json:"EmploymentSalaryType,omitempty"`

	// EmploymentSalaryTypeDescription: Salary type description
	EmploymentSalaryTypeDescription *string `json:"EmploymentSalaryTypeDescription,omitempty"`

	// EndDate: Salary record end date
	EndDate *types.Date `json:"EndDate,omitempty"`

	// FulltimeAmount: Salary when working fulltime
	FulltimeAmount *float64 `json:"FulltimeAmount,omitempty"`

	// HourlyWage: Hourly wage
	HourlyWage *float64 `json:"HourlyWage,omitempty"`

	// InternalRate: Internal rate for time &amp; billing or professional service user
	InternalRate *float64 `json:"InternalRate,omitempty"`

	// JobLevel: Employee job level in context of a wage scale
	JobLevel *int `json:"JobLevel,omitempty"`

	// Modified: Last modified date
	Modified *types.Date `json:"Modified,omitempty"`

	// Modifier: User ID of modifier
	Modifier *types.GUID `json:"Modifier,omitempty"`

	// ModifierFullName: Name of modifier
	ModifierFullName *string `json:"ModifierFullName,omitempty"`

	// ParttimeAmount: Salary when working parttime
	ParttimeAmount *float64 `json:"ParttimeAmount,omitempty"`

	// ParttimeFactor: Contract hours / Fulltime contract hours
	ParttimeFactor *float64 `json:"ParttimeFactor,omitempty"`

	// Scale: Employee wage scale
	Scale *string `json:"Scale,omitempty"`

	// Schedule: Employment schedule
	Schedule *types.GUID `json:"Schedule,omitempty"`

	// ScheduleCode: Employment schedule code
	ScheduleCode *string `json:"ScheduleCode,omitempty"`

	// ScheduleDescription: Description of employment schedule
	ScheduleDescription *string `json:"ScheduleDescription,omitempty"`

	// StartDate: Salary record start date
	StartDate *types.Date `json:"StartDate,omitempty"`
}

// List the EmploymentSalaries entities in the provided division.
// If all is true, all the paginated results are fetched; if false, list the first page.
func (s *EmploymentSalariesEndpoint) List(ctx context.Context, division int, all bool) ([]*EmploymentSalaries, error) {
	var entities []*EmploymentSalaries
	u, err := s.client.ResolvePathWithDivision("/api/v1/{division}/payroll/EmploymentSalaries?$select=*", division)
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
