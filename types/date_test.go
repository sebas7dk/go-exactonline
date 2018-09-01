// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package types

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDate_isSet(t *testing.T) {
	d := &Date{}

	if d.IsSet() {
		t.Error("Empty Date should return false for Date.IsSet")
	}
}

func TestDate_unMarshalJSON(t *testing.T) {
	b := []byte(`"/Date(1535718344017)/"`)

	var d Date
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("Unmarshalled Date should not return an err: %v", err)
	}

	if !d.IsSet() {
		t.Error("Unmarshalled Date should return true for Date.IsSet")
	}

	want := time.Date(2018, 8, 31, 12, 25, 44, 17000000, time.UTC)

	if !d.UTC().Equal(want) {
		t.Errorf("Date = %v, want %v", d, want)
	}
}

func TestDate_unMarshalJSONEmpty(t *testing.T) {
	tests := [][]byte{
		[]byte(`null`),
		[]byte(``),
		[]byte(`""`),
		[]byte(`/Date()/`),
	}

	for _, test := range tests {
		var d Date
		json.Unmarshal(test, &d)

		if d.IsSet() {
			t.Errorf("Unmarshalled Date with value: %s should return false for Date.IsSet", string(test))
		}
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	v := Date{time.Date(2018, 8, 31, 12, 25, 44, 17000000, time.UTC)}
	want := `"2018-08-31T12:25:44.017Z"`
	testJSONMarshal(t, v, want)

	v2 := time.Now()
	want2 := `"` + v2.Format(time.RFC3339Nano) + `"`
	testJSONMarshal(t, v2, want2)
}

// Helper function to test that a value is marshalled to JSON as expected.
func testJSONMarshal(t *testing.T, v interface{}, want string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	if err != nil {
		t.Errorf("String is not valid json: %s", want)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}
}

func testJSONUnmarshal(t *testing.T, v interface{}, want string) {
	// now go the other direction and make sure things unmarshal as expected
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(want), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v", want)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, want %s", want, u, v)
	}
}
