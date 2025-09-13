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

func TestMySum(t *testing.T) {
	actual := mySum(5, 6, 7)
	expected := 18
	if actual != expected {
		t.Error("Expected", expected, " but Got", actual)
	}
}

func TestMySumWrong(t *testing.T) {
	actual := mySumWrong(5, 6, 7)
	expected := 18
	if actual != expected {
		t.Error("Expected", expected, " but Got", actual)
	}
}
