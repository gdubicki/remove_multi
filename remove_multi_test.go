package main

import (
	"testing"
	"log"
	"strings"
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
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

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveFirstLineDoesntMatch(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
	}
	toRemoveLines := []string{
		"Line 2",
		"Line that doesn't match",
		"Line 4",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
	}

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveSomeRegexp(t *testing.T) {
	targetLines := string_to_lines(`
some lines...

foo::bar::foo: barbar
abc::def::zxy:
  something_else:
    paths:
      - /opt/tomcat/path/logs/filer_whatever.log
    fields:
      datacenter: dc
      type: filer_whatever
    multiline:
      pattern: ^(ERROR|INFO|WAR|TRACE|NOTICE|ALERT|CRIT|DEBUG)
      negate: true
      match: after

some other lines`)
	toRemoveLines := string_to_lines(`abc::def::zxy:
  .*_else:
    paths:
      - /opt/tomcat/path/logs/.*_whatever.log
    fields:
      datacenter:.*
      type: .*
    multiline:
      pattern: .*
      negate: .*
      match: .*`)

	expectedResultLines := string_to_lines(`
some lines...

foo::bar::foo: barbar

some other lines`)

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
}

func TestRemoveSomeRegexpDontMatchAll(t *testing.T) {
	targetLines := string_to_lines(`
some lines...

foo::bar::foo: barbar
abc::def::zxy:
  something_not-matching:
    paths:
      - /opt/tomcat/path/logs/filer_whatever.log
    fields:
      datacenter: dc
      type: filer_whatever
    multiline:
      pattern: ^(ERROR|INFO|WAR|TRACE|NOTICE|ALERT|CRIT|DEBUG)
      negate: true
      match: after

some other lines`)

	toRemoveLines := string_to_lines(`abc::def::zxy:
  .+_else:
    paths:
      - /opt/tomcat/path/logs/.*_whatever.log
    fields:
      datacenter:.*
      type: .*
    multiline:
      pattern: .*
      negate: .*
      match: .*`)

	expectedResultLines := targetLines

	remove_helper(targetLines, toRemoveLines, expectedResultLines, t)
}


func TestReplaceSingleLine(t *testing.T) {
	targetLines := []string {
		"Line 1",
		"Line 2",
		"Line 3",
	}
	toRemoveLines := []string{
		"Line 2",
	}
	toReplaceLines := []string{
		"Line ABC",
	}

	expectedResultLines := []string {
		"Line 1",
		"Line ABC",
		"Line 3",
	}

	replace_helper(targetLines, toRemoveLines, toReplaceLines, expectedResultLines, t)
}


func string_to_lines(aString string) []string {
	return strings.Split(aString, "\n")
}

func remove_helper(targetLines []string, toRemoveLines []string, expectedResultLines []string, t *testing.T) {
	result, err := remove(targetLines, toRemoveLines, nil)

	if err != nil {
		log.Fatal(err)
	}

	if ! testEq(result, expectedResultLines) {
		t.Errorf(
			"For '%s' - '%s'\nexpected '%s'\ngot '%s'\n", targetLines, toRemoveLines, expectedResultLines, result,
		)
	}
}

func replace_helper(targetLines []string, toRemoveLines []string, toReplaceLines[]string, expectedResultLines []string, t *testing.T) {
	result, err := remove(targetLines, toRemoveLines, toReplaceLines)

	if err != nil {
		log.Fatal(err)
	}

	if ! testEq(result, expectedResultLines) {
		t.Errorf(
			"For '%s' + ('%s' -> '%s')\nexpected '%s'\ngot '%s'\n", targetLines, toRemoveLines, toReplaceLines, expectedResultLines, result,
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