package data

import (
	"reflect"
	"testing"
)

// TestPosition_Copy tests the Copy method of the Position struct.
// It verifies that the copied position is equal to the original position,
// and that modifying the original position does not affect the copied position.
func TestPosition_Copy(t *testing.T) {
	original := Position{
		Timestamp:  1629876543,
		Latitude:   37.7749,
		Longitude:  -122.4194,
		Altitude_M: 10.5,
		Speed_kMh:  50.0,
	}

	copy := original.Copy()

	if !reflect.DeepEqual(original, copy) {
		t.Errorf("Expected copied position to be equal to original, got %v", copy)
	}

	// Modifying the original position should not affect the copied position
	original.Timestamp = 1629876544
	original.Latitude = 37.7748
	original.Longitude = -122.4193
	original.Altitude_M = 11.5
	original.Speed_kMh = 55.0

	if reflect.DeepEqual(original, copy) {
		t.Errorf("Expected copied position to be different from original, got %v", copy)
	}
}
func TestPosition_SQLCreateQuery(t *testing.T) {
	p := Position{
		Timestamp:  1629876543,
		Latitude:   37.7749,
		Longitude:  -122.4194,
		Altitude_M: 10.5,
		Speed_kMh:  50.0,
		Course:     90.0,
	}

	expectedQuery := "CREATE TABLE IF NOT EXISTS positions (timestamp_unix INTEGER PRIMARY KEY, latitude REAL, longitude REAL, altitude_m REAL, speed_kmh REAL, course_degrees REAL)"

	query := p.SQLCreateQuery()

	if query != expectedQuery {
		t.Errorf("Expected query: %s, got: %s", expectedQuery, query)
	}
}

func TestPosition_SQLInsertQuery(t *testing.T) {
	p := Position{
		Timestamp:  1629876543,
		Latitude:   37.7749,
		Longitude:  -122.4194,
		Altitude_M: 10.5,
		Speed_kMh:  50.0,
		Course:     90.0,
	}

	expectedQuery := "INSERT INTO positions (timestamp_unix, latitude, longitude, altitude_m, speed_kmh, course_degrees) VALUES (1629876543, 37.7749, -122.4194, 10.5, 50, 90)"

	query := p.SQLInsertQuery()

	if query != expectedQuery {
		t.Errorf("Expected query: %s, got: %s", expectedQuery, query)
	}
}
