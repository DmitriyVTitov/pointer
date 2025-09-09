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

	var v any // nil
	if got := *To(v); got != nil {
		t.Errorf("got %v want %v", got, nil)
	}
}

func TestVal(t *testing.T) {
	b := true
	if got := Val(&b); got != b {
		t.Errorf("got %v want %v", got, b)
	}

	s := "string"
	if got := Val(&s); got != s {
		t.Errorf("got %v want %v", got, s)
	}

	var v any // nil
	if got := Val(&v); got != nil {
		t.Errorf("got %v want %v", got, nil)
	}
}
