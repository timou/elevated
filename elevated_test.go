package elevated

import (
	"testing"
)

func TestIsElevated(t *testing.T) {
	// Test that we don't crash.
	got := IsElevated()
	_ = got
}