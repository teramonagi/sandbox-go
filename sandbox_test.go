package main

import "testing"

func TestHello(t *testing.T) {
	if hello() != "Hello, world!" {
		t.Errorf("Not hello world")
	}
}
