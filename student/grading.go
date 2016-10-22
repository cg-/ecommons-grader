package student

import (
	"archive/zip"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cg-/ecommons-grader/assignment"
)

func (s *Student) prepareFileToGrade(filetype string) (string, error) {
	var filename os.File
	var fileInZip *zip.File

	foundFile := false
	for i := range s.Files {
		if strings.Contains(s.Files[i].Name, filetype) {
			fileInZip = s.Files[i]
			foundFile = true
		}
	}
	if foundFile == false {
		return "", fmt.Errorf("Couldn't find an appropriate file to grade.")
	}
	filepath.Join(os.TempDir(), f.Name())
}

func (s *Student) Grade(a *assignment.Assignment) {
	for gradeCommand, questionArray := range a.Content {
		var wg sync.WaitGroup
		wg.Add(1)
		file, err := s.prepareFileToGrade(gradeCommand.Filetype)
		go func(cmd string) {
			defer wg.Done()
			exec.Command(cmd)
		}(gradeCommand.Command + " " + file)

		wg.Wait()
	}
}
