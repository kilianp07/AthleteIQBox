package transmitter

import (
	"testing"
)

func TestFactory(t *testing.T) {
	testCases := []struct {
		name        string
		serviceType string
		wantErr     bool
	}{
		{"Valid service type", "position", false},
		{"Invalid service type", "invalidService", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := factory(tc.serviceType, nil)
			if tc.wantErr {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}
