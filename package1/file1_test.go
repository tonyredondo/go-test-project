package package1

import "testing"

func TestValue(t *testing.T) {
	if value() != "value" {
		t.Errorf("Expected 'value', got '%s'", value())
	}
}
