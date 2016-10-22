package student

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

// GRADEFILE is the name of the file containing the grades in the zip file.
const GRADEFILE string = "grades.csv"

// CsvInfo holds the info about the student from the Csv File
type CsvInfo struct {
	DisplayID string
	ID        string
	LastName  string
	FirstName string
	Grade     string
	CsvLine   int
}

// Student is a structure that will hold pointers to the various files inside
// the downloaded zip file. We will use these while we process the data.
type Student struct {
	Name        string
	CommentFile *zip.File
	Files       []*zip.File
	Csv         CsvInfo
}

// ParseWorkDir parses a work directory and returns a student map
func ParseWorkDir(path string) (map[string]*Student, error) {
	var returnMap = make(map[string]*Student)

	log.Fatalf("Parsing Work Dir not yet implemented.\n")

	return returnMap, nil
}

// ParseGradeFile parses a zip and generates a student map
func ParseGradeFile(z *zip.ReadCloser) (map[string]*Student, map[string][]string, error) {
	var returnMap = make(map[string]*Student)
	var returnIgnoredLinesMap = make(map[string][]string)
	// Run through and put together the returnMap
	labDir := strings.Split(z.File[0].Name, "/")[0] + "/"
	var gradeCsv io.ReadCloser
	var err error

	for i := range z.File {
		if z.File[i].Name == labDir+GRADEFILE {
			log.Printf("Found the grades file [%s]", z.File[i].Name)
			gradeCsv, err = z.File[i].Open()
			if err != nil {
				log.Fatalf("Couldn't open the grade csv file: %s", err.Error())
			}
		} else {
			studentNameSplit := strings.Split(z.File[i].Name, "/")
			if returnMap[studentNameSplit[1]] == nil {
				returnMap[studentNameSplit[1]] = &Student{
					Name: studentNameSplit[1],
				}
			}
			returnMap[studentNameSplit[1]].Files = append(returnMap[studentNameSplit[1]].Files, z.File[i])
		}
	}

	// Open and read the grade file.
	r := csv.NewReader(bufio.NewReader(gradeCsv))
	csvFileLine := -1
	for {
		csvFileLine++
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		// Skip over records that don't matter (headers, blank, etc...)
		if len(record) != 5 || record[1] == "ID" {
			log.Printf("Ignoring %s", record)
			returnIgnoredLinesMap[strconv.Itoa(csvFileLine)] = record
			continue
		}

		newCsv := CsvInfo{
			DisplayID: record[0],
			ID:        record[1],
			LastName:  record[2],
			FirstName: record[3],
			Grade:     record[4],
			CsvLine:   csvFileLine,
		}

		// put the csv info on the student struct
		matchedName := false
		for k, v := range returnMap {
			if strings.Contains(k, newCsv.ID) {
				v.Csv = newCsv
				matchedName = true
				break
			}
		}
		if matchedName == false {
			log.Fatalf("Couldn't match %s", newCsv.ID)
		}
	}
	return returnMap, returnIgnoredLinesMap, nil
}
