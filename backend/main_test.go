package main

import "testing"

func TestSum(t *testing.T) {
	want := 5

	got := sum(2, 3)

	if got != want {
		t.Fatalf("sum(2, 3) = %v, want %v", got, want)
	}
}
