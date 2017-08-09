package main

import (
	"testing"
	"log"
)

func TestRemoveSingleLine(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
	}
	toRemoveLines := []string{
		"Line 2",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line 3",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveSingleLineFromBeginning(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
	}
	toRemoveLines := []string{
		"Line 1",
	}

	expectedResultLines := []string {
		"Line 2",
		"Line 3",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveSingleLineFromEnd(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
	}
	toRemoveLines := []string{
		"Line 3",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line 2",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveMoreLines(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
	}
	toRemoveLines := []string{
		"Line 2",
		"Line 3",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line 4",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveMoreLinesFromBeginning(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
	}
	toRemoveLines := []string{
		"Line 1",
		"Line 2",
	}

	expectedResultLines := []string {
		"Line 3",
		"Line 4",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveMoreLinesFromEnd(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
	}
	toRemoveLines := []string{
		"Line 3",
		"Line 4",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line 2",
	}

	helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func helper(targetLines []string, toRemoveLines []string, expectedResultLines []string, t *testing.T) {
	result, err := remove(targetLines, toRemoveLines)

	if err != nil {
		log.Fatal(err)
	}

	if ! testEq(result, expectedResultLines) {
		t.Errorf(
			"For '%s' - '%s'\nexpected '%s'\ngot '%s'\n", targetLines, toRemoveLines, expectedResultLines, result,
		)
	}
}
func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}