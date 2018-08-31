// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package exactonline

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestCRMQuotationLinesService_List_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("/api/v1/{division}/crm/QuotationLines?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("/api/v1/{division}/crm/QuotationLines?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in CRMQuotationLinesService.List returned error: %v, with url /api/v1/{division}/crm/QuotationLines", e)
	}

	g := NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
		}
	})

	/* g := NewGUID()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "` + u2.String() + `", "results": [{ "ID": "` + g.String() + `"}]}}`)
		}
	}) */

	entities, err := client.CRMQuotationLines.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("CRMQuotationLinesService.List returned error: %v", err)
	}

	want := []*CRMQuotationLines{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CRMQuotationLinesService.List returned %+v, want %+v", entities, want)
	}
}

func TestCRMQuotationLinesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := client.ResolvePathWithDivision("/api/v1/{division}/crm/QuotationLines?$select=*", 0)
	u2, e := client.ResolvePathWithDivision("/api/v1/{division}/crm/QuotationLines?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("client.ResolvePathWithDivision in CRMQuotationLinesService.List returned error: %v, with url /api/v1/{division}/crm/QuotationLines", e)
	}

	g := NewGUID()
	gs := g.String()
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "ID": "`+gs+`"}]}}`)
	})

	entities, err := client.CRMQuotationLines.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("CRMQuotationLinesService.List returned error: %v", err)
	}

	want := []*CRMQuotationLines{{ID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("CRMQuotationLinesService.List returned %+v, want %+v", entities, want)
	}
}
