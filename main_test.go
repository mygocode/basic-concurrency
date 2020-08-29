package main

import (
	"testing"
)

func TestWithPositiveInteger(t *testing.T) {

	var workers int = 2
	var jobs int = 20
	var expected int = 20

	if output := concurrentFunc(workers, jobs); output != expected {
		t.Error("Test Failed: {} expected, recieved: {}", expected, output)
	}
}

func TestWithNegativeInteger(t *testing.T) {

	var workers int = -2
	var jobs int = 20
	var expected int = -1

	if output := concurrentFunc(workers, jobs); output != expected {
		t.Error("Test Failed: {} expected, recieved: {}", expected, output)
	}
}
