package domain

import (
	"encoding/json"
	"testing"
)

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error ocurred: %s", err.Error())
	}
}

func assertInterfaceEqual(t *testing.T, expected interface{}, got interface{}) {
	t.Helper()
	expectedJson, _ := json.Marshal(expected)
	gotJson, _ := json.Marshal(got)

	if string(expectedJson) != string(gotJson) {
		t.Errorf("expected:\n%s\ngot:\n%s", expectedJson, gotJson)
	}
}
