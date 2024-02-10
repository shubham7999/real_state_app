package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	// Capture stdout to check if the expected message is printed
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	InitDatabase()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	expectedMessage := "Database connected successfully"
	if !strings.Contains(string(out), expectedMessage) {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, out)
	}
}

func TestGenerateID(t *testing.T) {
	// Initialize the global ID to 0 before running the test
	globalId = 0

	// Generate IDs and check if they are unique
	idSet := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		id := GenerateID()
		if _, exists := idSet[id]; exists {
			t.Errorf("Generated duplicate ID: %s", id)
		}
		idSet[id] = struct{}{}
	}

	// Check if the global ID is incremented correctly
	expectedGlobalID := 100
	if globalId != expectedGlobalID {
		t.Errorf("Expected global ID to be %d, got %d", expectedGlobalID, globalId)
	}
}
