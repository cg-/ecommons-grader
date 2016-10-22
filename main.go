package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cg-/ecommons-grader/student"
	"github.com/davecgh/go-spew/spew"
)

// OriginalStudentMap is the map of data we need to process.
var OriginalStudentMap map[string]*student.Student

// OutputStudentMap is the map of data we have already processed.
var OutputStudentMap map[string]*student.Student

// Command Line Arguments
var inputZipFlag = flag.String("input", "bulk_download.zip", "Path to the eCommons ZIP file.")
var workDir = flag.String("work", ".ecommons-work", "Path to the work directory.")

//outputZipFlag := flag.String("output", "output.zip", "Path to the output ZIP file.")

func checkArguments() {
	flag.Parse()

	_, err := ioutil.ReadDir(*workDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(*workDir, 0755)
		} else {
			log.Fatalf("Err opening output folder [%s]", err.Error())
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("Output Directory is already specified. What do you want to do?")
			fmt.Println("1. Continue Working")
			fmt.Println("2. Delete it")
			fmt.Println("3. Quit")
			text, _ := reader.ReadString('\n')
			val := strings.TrimSpace(strings.ToLower(text))
			if val == "1" {
				student.ParseWorkDir(*workDir)
				break
			} else if val == "2" {
				fmt.Println("ARE YOU SURE? (y/n)")
				confirmText, _ := reader.ReadString('\n')
				confirmVal := strings.TrimSpace(strings.ToLower(confirmText))
				if confirmVal == "y" {
					os.RemoveAll(*workDir)
					os.Mkdir(*workDir, 0755)
					break
				} else {
					log.Fatalf("Will not proceed. Please specify a different output directory.")
				}
			} else if val == "3" {
				os.Exit(0)
			}
		}
	}
}

func main() {
	spew.Dump(OriginalStudentMap)
	checkArguments()

	// Open the original input file
	z, err := zip.OpenReader(*inputZipFlag)
	if err != nil {
		log.Fatalf("Err opening zip file \"" + *inputZipFlag + "\"" + err.Error())
	}
	defer z.Close()

	// Parse the original input file
	s, i, err := student.ParseGradeFile(z)
	spew.Dump(i)

	// Now we start grading.
	for k, _ := range s {
		log.Printf("%s", k)
	}
}
