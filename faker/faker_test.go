package faker

import (
	"testing"
)

func TestName(t *testing.T) {
	name := Name()
	if name == "" {
		t.Error("Expected a name, got an empty string")
	}
}
