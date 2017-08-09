package main

import (
	"bufio"
	"os"
	"log"
)

// TODO: use something like Python's argparse for a nice cli
func main() {

	targetFile := os.Args[1]
	toRemoveFile := os.Args[2]

	targetFileLines, err := readLines(targetFile)
	if err != nil {
		log.Fatal(err)
	}

	toRemoveFileLines, err := readLines(toRemoveFile)
	if err != nil {
		log.Fatal(err)
	}

	resultFileLines, err := remove(targetFileLines, toRemoveFileLines)
	if err != nil {
		log.Fatal(err)
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
