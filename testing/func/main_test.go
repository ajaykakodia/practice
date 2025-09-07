package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	actual := add(5, 6)
	expected := 11
	if actual != expected {
		log.Fatalf("We are expecting %d but got %d", expected, actual)
	}
}
