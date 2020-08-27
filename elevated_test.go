package elevated

import (
	"testing"
)

func TestIsElevated(t *testing.T) {
	// The only real test is that it doesn't panic on Windows.
	got := IsElevated()
	_ = got
}