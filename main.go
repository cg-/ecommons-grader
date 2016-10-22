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
)

// GRADEFILE is the name of the file containing the grades in the zip file.
const GRADEFILE string = "grades.csv"

// StudentMap is a map of Student structs that will hold all the data while
// we process it.
var StudentMap = make(map[string]*Student)

// Student is a structure that will hold pointers to the various files inside
// the downloaded zip file. We will use these while we process the data.
type Student struct {
	Name        string
	ID          string
	CommentFile *zip.File
	Grade       int
	Files       []*zip.File
}

func main() {
	// Parse options
	inputZipFlag := flag.String("input", "bulk_download.zip", "Path to the eCommons ZIP file.")
	outputDir := flag.String("output", "output", "Path to the output directory.")
	flag.Parse()

	_, err := ioutil.ReadDir(*outputDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(*outputDir, 0755)
		} else {
			log.Fatalf("Err opening output folder [%s]", err.Error())
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("Output Directory is already specified. Do you want to delete it? (y/n)")
			text, _ := reader.ReadString('\n')
			val := strings.TrimSpace(strings.ToLower(text))
			if val == "y" {
				fmt.Println("ARE YOU SURE? (y/n)")
				confirmText, _ := reader.ReadString('\n')
				confirmVal := strings.TrimSpace(strings.ToLower(confirmText))
				if confirmVal == "y" {
					os.RemoveAll(*outputDir)
					os.Mkdir(*outputDir, 0755)
					break
				} else {
					log.Fatalf("Will not proceed. Please specify a different output directory.")
				}
			} else if val == "n" {
				log.Fatalf("Will not proceed. Please specify a different output directory.")
			}
		}
	}

	z, err := zip.OpenReader(*inputZipFlag)
	if err != nil {
		log.Fatalf("Err opening zip file \"" + *inputZipFlag + "\"" + err.Error())
	}
	defer z.Close()

	labDir := strings.Split(z.File[0].Name, "/")[0] + "/"

	for i := range z.File {
		if z.File[i].Name == labDir+GRADEFILE {
			log.Printf("Found the grades file [%s]", z.File[i].Name)
		} else {
			studentNameSplit := strings.Split(z.File[i].Name, "/")
			if StudentMap[studentNameSplit[1]] == nil {
				StudentMap[studentNameSplit[1]] = &Student{}
			}
			StudentMap[studentNameSplit[1]].Files = append(StudentMap[studentNameSplit[1]].Files, z.File[i])
		}
	}

	i := 0
	for k, _ := range StudentMap {
		i++
		fmt.Println(k)
	}
	fmt.Printf("Total students: %d\n", i)
}
