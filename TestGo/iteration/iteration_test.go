package iteration

import (
	"testing"
)

func TestIteration(t *testing.T) {
	result := Result("a")
	expected := "aaaaa"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
