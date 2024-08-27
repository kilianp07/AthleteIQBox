package sqlite

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kilianp07/AthleteIQBox/data"
)

func TestRecorder_Configure(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "recorder_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Create a mock position channel
	positionCh := make(chan data.Position)

	// Create a mock configuration with the temporary directory
	conf := configuration{
		Filepath: tempDir,
	}

	// Create a new Recorder instance
	r := &Recorder{
		conf:       conf,
		positionCh: positionCh,
	}

	// Call the Configure method
	if err := r.Configure(positionCh); err != nil {
		t.Fatalf("Unexpected error during Configure: %v", err)
	}

	// Check that the database was created
	expectedDBFile := fmt.Sprintf("%s%s.db", r.conf.Filepath, time.Now().Format(time.RFC3339))
	if _, err := os.Stat(expectedDBFile); os.IsNotExist(err) {
		t.Fatalf("Expected database file was not created: %v", expectedDBFile)
	}

	// Check if the database connection is open
	if err := r.db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// Check if the table was created
	rows, err := r.db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='positions'")
	if err != nil {
		t.Fatalf("Failed to query tables: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Fatalf("Expected table 'positions' was not created")
	}

	// Close the database connection at the end of the test
	if err := r.db.Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}

}
