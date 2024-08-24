package utils

import (
	"reflect"
	"testing"

	"github.com/go-viper/mapstructure/v2"
)

func TestNewDecoder(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name           string
		output         interface{}
		opts           []DecoderOption
		expectedResult *mapstructure.Decoder
		expectedErr    bool
	}{
		{
			name:           "Valid output and options",
			output:         &TestData{},
			opts:           []DecoderOption{WithTagName("yaml")},
			expectedResult: &mapstructure.Decoder{},
			expectedErr:    false,
		},
		{
			name:           "Invalid output (not a pointer)",
			output:         TestData{},
			opts:           []DecoderOption{},
			expectedResult: nil,
			expectedErr:    true,
		},
		{
			name:           "Valid output and default options",
			output:         &TestData{},
			opts:           []DecoderOption{},
			expectedResult: &mapstructure.Decoder{},
			expectedErr:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			decoder, err := NewDecoder(tc.output, tc.opts...)
			if tc.expectedErr {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got '%v'", err)
				} else if !reflect.DeepEqual(reflect.TypeOf(decoder), reflect.TypeOf(tc.expectedResult)) {
					t.Errorf("Expected decoder type '%v', got '%v'", reflect.TypeOf(tc.expectedResult), reflect.TypeOf(decoder))
				}
			}
		})
	}
}
