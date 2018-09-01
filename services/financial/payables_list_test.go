// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package financial

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestPayablesListEndpoint_List_all(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList?$select=*", 0)
	u2, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PayablesListEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/PayablesList", e)
	}

	g := int64(100)
	gs := strconv.Itoa(int(g))
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		if r.URL.Query().Get("$skiptoken") != "" {
			fmt.Fprint(w, `{ "d": { "__next": "", "results": []}}`)
		} else {
			fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "HID": `+gs+`}]}}`)
		}
	})

	entities, err := s.PayablesList.List(context.Background(), 0, true)
	if err != nil {
		t.Errorf("PayablesListEndpoint.List returned error: %v", err)
	}

	want := []*PayablesList{{HID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PayablesListEndpoint.List returned %+v, want %+v", entities, want)
	}
}

func TestPayablesListEndpoint_List(t *testing.T) {
	s, mux, _, teardown := setup()
	defer teardown()

	acceptHeaders := []string{"application/json"}

	u, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList?$select=*", 0)
	u2, e := s.client.ResolvePathWithDivision("/api/v1/{division}/read/financial/PayablesList?$skiptoken=foo", 0)
	if e != nil {
		t.Errorf("s.client.ResolvePathWithDivision in PayablesListEndpoint.List returned error: %v, with url /api/v1/{division}/read/financial/PayablesList", e)
	}

	g := int64(100)
	gs := strconv.Itoa(int(g))
	mux.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(acceptHeaders, ", "))
		fmt.Fprint(w, `{ "d": { "__next": "`+u2.String()+`", "results": [{ "HID": `+gs+`}]}}`)
	})

	entities, err := s.PayablesList.List(context.Background(), 0, false)
	if err != nil {
		t.Errorf("PayablesListEndpoint.List returned error: %v", err)
	}

	want := []*PayablesList{{HID: &g}}
	if !reflect.DeepEqual(entities, want) {
		t.Errorf("PayablesListEndpoint.List returned %+v, want %+v", entities, want)
	}
}