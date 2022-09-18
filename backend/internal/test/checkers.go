package test

import (
	"reflect"
	"testing"
)

// IsNil asserts that err is nil.
func IsNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
}

// NotNil asserts that err is not nil.
func NotNil(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

// DeepEqual asserts that want DeepEquals to got.
func DeepEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %#v, got: %#v", want, got)
	}
}
