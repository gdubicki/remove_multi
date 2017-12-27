package main

import (
	"log"
	"regexp"
	"fmt"
)


func remove(targetLines []string, toRemoveLines []string, toReplaceLines []string) ([] string, error) {
	// Let's start with naive implementation that works and more forward from here

	var toReturn []string

	if len(toRemoveLines) > len(targetLines) {
		log.Fatal("You need less to remove lines than target lines!")
	}

	if toReplaceLines != nil && len(toRemoveLines) != len(toReplaceLines) {
		log.Fatal("You need same number of lines in to_replace as in to_remove files!")
	}

	compiledLines, err := compileLines(toRemoveLines)

	if err != nil {
		return nil, err
	}

	var allMatch bool

	var i = 0

	// Use sliding window
	for ; i < len(targetLines) - len(toRemoveLines) + 1; {

		allMatch = true

		// Test if all lines [i, len(compiledLines)] match all compiledLines
		for j := 0; j < len(compiledLines); j++ {
			singleMatches := compiledLines[j].FindStringSubmatch(targetLines[i + j]) != nil

			if ! singleMatches {
				allMatch = false
				break
			}
		}

		if allMatch {
			if toReplaceLines != nil {
				toReturn = append(toReturn, toReplaceLines...)
			}
			i += len(toRemoveLines)
		} else {
			toReturn = append(toReturn, targetLines[i])
			i++
		}

	}

	// If the are any lines remaining then add them
	if i < len(targetLines) {
		restOfLines := targetLines[i:]
		toReturn = append(toReturn, restOfLines...)
	}

	return toReturn, nil
}

// Compile regexps in toRemoveLines
func compileLines(toRemoveLines []string) ([]*regexp.Regexp, error) {
	var toReturn []*regexp.Regexp

	for j := 0; j < len(toRemoveLines); j++ {

		// Wrap line with '^' to '$' to make sure we match whole line only
		strRegexp := fmt.Sprintf("^%v$", toRemoveLines[j])
		compiledRegexp, err := regexp.Compile(strRegexp)

		if err != nil {
			return nil, err
		}

		toReturn = append(toReturn, compiledRegexp)

	}

	return toReturn, nil
}