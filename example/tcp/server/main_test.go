package main

import "testing"

func TestPong(t *testing.T) {
	if err := Pong(":8080"); err != nil {
		t.Fatal(err)
	}
}
