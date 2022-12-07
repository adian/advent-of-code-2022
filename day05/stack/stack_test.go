package stack

import (
	"reflect"
	"testing"
)

func Test_stack(t *testing.T) {
	s := New[string]()

	got, ok := s.TakeOne()
	if ok {
		t.Errorf("ok should return false, got %v", ok)
	}
	if got != "" {
		t.Errorf("got %v, want empty value", got)
	}

	want := "asdf"
	s.Put(want)
	got, ok = s.TakeOne()
	if !ok {
		t.Errorf("ok should return true, got %v", ok)
	}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got, ok = s.TakeOne()
	if ok {
		t.Errorf("ok should return false, got %v", ok)
	}
	if got != "" {
		t.Errorf("got %v, want empty value", got)
	}
}

func errorF[T any](t *testing.T, got T, want T) {
	t.Errorf("got %#v, want %#v", got, want)
}

func Test_stack_take(t *testing.T) {
	s := New[string]()

	s.PutAll("a", "b", "c", "d")
	got := s.Take(2)
	want := []string{"c", "d"}
	if !reflect.DeepEqual(got, want) {
		errorF(t, got, want)
	}
}
