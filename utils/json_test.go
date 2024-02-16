package utils

import (
	"reflect"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	testCases := []struct {
		name         string
		path         string
		target       interface{}
		expectedErr  bool
		expectedData interface{}
	}{
		{
			name:         "Valid JSON file",
			path:         "../tests/json/valid.json",
			target:       &struct{ Name string }{},
			expectedErr:  false,
			expectedData: &struct{ Name string }{Name: "John Doe"},
		},
		{
			name:         "Invalid JSON file",
			path:         "../tests/json/invalid.json",
			target:       &struct{ Age int }{},
			expectedErr:  true,
			expectedData: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ReadJSONFile(tc.path, tc.target)
			if tc.expectedErr {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got '%v'", err)
				} else if !reflect.DeepEqual(tc.target, tc.expectedData) {
					t.Errorf("Expected data '%v', got '%v'", tc.expectedData, tc.target)
				}
			}
		})
	}
}
