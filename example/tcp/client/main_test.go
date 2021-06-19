package main

import "testing"

func TestPing(t *testing.T) {
	if err := Ping(":8080"); err != nil {
		t.Fatal(err)
	}
}
