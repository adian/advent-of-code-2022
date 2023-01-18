package asserts

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Helper()
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func DeepEqual[T any](t *testing.T, got, want T) {
	if !reflect.DeepEqual(got, want) {
		t.Helper()
		t.Errorf("got %#v, want %#v", got, want)
	}
}
