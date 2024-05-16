package main

import (
	l "dll-tdd/double-linked-list"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected [%+v], got: [%+v]", expected, actual)
	}
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected to be nil, found: [%+v]", actual)
	}
}

func TestNewList_ReturnsEmptyList(t *testing.T) {
	myList := l.New()
	assertEqual(t, myList.Length, 0)
	assertNil(t, myList.Head)
}

func TestNewList_ReturnsEmptyLength(t *testing.T) {
	myList := l.New()
	assertEqual(t, myList.Len(), 0)
}
