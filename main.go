package main

import (
	"bufio"
	"os"
	"log"
)

// TODO: use something like Python's argparse for a nice cli
func main() {
	var resultFileLines []string

	operation := os.Args[1] // "remove" or "replace"

	targetFile := os.Args[2]
	toRemoveFile := os.Args[3]

	targetFileLines, err := readLines(targetFile)
	if err != nil {
		log.Fatal(err)
	}

	toRemoveOrReplaceFileLines, err := readLines(toRemoveFile)
	if err != nil {
		log.Fatal(err)
	}

	if operation == "remove" {
		resultFileLines, err = remove(targetFileLines, toRemoveOrReplaceFileLines, nil)
	} else if operation == "replace" {
		toReplaceFile := os.Args[4]

		toReplaceFileLines, err := readLines(toReplaceFile)
		if err != nil {
			log.Fatal(err)
		}

		resultFileLines, err = remove(targetFileLines, toRemoveOrReplaceFileLines, toReplaceFileLines)
	}

	writeLines(targetFile, resultFileLines)
	if err != nil {
		log.Fatal(err)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// write lines to file
func writeLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
