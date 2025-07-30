package pointer

import (
	"testing"
)

func TestTo(t *testing.T) {
	b := true
	if got := *To(b); got != b {
		t.Errorf("got %v want %v", got, b)
	}

	s := "string"
	if got := *To(s); got != s {
		t.Errorf("got %v want %v", got, s)
	}

	var v any
	if got := *To(v); got != nil {
		t.Errorf("got %v want %v", got, nil)
	}
}
