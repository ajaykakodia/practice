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

func TestMySumTableTest(t *testing.T) {
	type test struct {
		data           []int
		expectedOutput int
	}

	tests := []test{
		{data: []int{21, 21}, expectedOutput: 42},
		{data: []int{10, 10, 22}, expectedOutput: 42},
		{data: []int{10, 10, 10, 12}, expectedOutput: 42},
		{data: []int{5, 5, 5, 5, 6, 6, 10}, expectedOutput: 42},
		{data: []int{-1, -1, 2}, expectedOutput: 0},
	}

	for _, val := range tests {
		result := mySum(val.data...)
		if result != val.expectedOutput {
			t.Error("Expected", val.expectedOutput, "but Got", result)
		}
	}
}

func TestMySumWrongTableTest(t *testing.T) {
	type test struct {
		data           []int
		expectedOutput int
	}

	tests := []test{
		{data: []int{21, 21}, expectedOutput: 42},
		{data: []int{10, 10, 22}, expectedOutput: 42},
		{data: []int{10, 10, 10, 12}, expectedOutput: 42},
		{data: []int{5, 5, 5, 5, 6, 6, 10}, expectedOutput: 42},
		{data: []int{-1, -1, 2}, expectedOutput: 0},
	}

	for _, val := range tests {
		result := mySumWrong(val.data...)
		if result != val.expectedOutput {
			t.Error("Expected", val.expectedOutput, "but Got", result)
		}
	}
}
