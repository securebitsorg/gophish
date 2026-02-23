package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	actual := Hello() // Assuming Hello() is the function you want to test
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}