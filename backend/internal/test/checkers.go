package test

import (
	"reflect"
	"testing"
)

func IsNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
}

func NotNil(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
}

func DeepEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %#v, got: %#v", want, got)
	}
}
