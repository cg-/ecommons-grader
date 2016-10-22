package assignment

import (
	"bufio"
	"log"
	"os"
)

// A GradeCommand represents a single command that needs to be executed to
// grade an assignment
type GradeCommand struct {
	Command  string
	Filetype string
}

// Question is a single question on an assignment
type Question struct {
	Question     string
	Points       int
	CommonErrors []map[string]int
}

// Assignment is a full assignment to be graded
type Assignment struct {
	Name        string
	TotalPoints int
	Content     map[GradeCommand][]Question
}

// ParseAssignmentFile will parse an input file and generate an Assignment from it
func ParseAssignmentFile(path string) (*Assignment, error) {
	returnAssignment := Assignment{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		log.Printf("Parsing: [%s]", currentLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &returnAssignment, nil
}
