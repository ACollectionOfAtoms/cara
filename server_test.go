package main

import (
	"bytes"
	"testing"
)

func TestStringsFromIOReader(t *testing.T) {
	t.Log("testing stringFromIOReader...")
	s := "A byte array?"
	b := []byte(s)
	r := bytes.NewReader(b)
	newString := stringFromIOReader(r)
	if newString != s {
		t.Errorf("Expected %s, but got %s!", s, newString)
	}
}
