package osutil

import (
	"github.com/karrick/godirwalk"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
	"strings"
)

// BinExists checks binary name in the user's PATH
func BinExists(binname string) bool {
	if _, err := PrintBinaryPath(binname); err != nil {
		return false
	}
	return true
}

// PrintBinaryPath prints binary location in the user's PATH
func PrintBinaryPath(binname string) (string, error) {
	return exec.LookPath(binname)
}

// WalkPath gets all the directory names in the specified directory
// returns all the directory paths that contains a Makefile.
func WalkPath(dir string) (dirs map[string]string, err error) {
	dirs = map[string]string{}
	log.Printf("Scanning : %s\n", dir)
	var allDirs []string
	if err := godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(pathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				allDirs = append(allDirs, filepath.FromSlash(pathname))
			}
			return nil
		},
		ErrorCallback: func(pathname string, err error) godirwalk.ErrorAction {
			log.Errorf("ERROR: %v\n", err)
			return godirwalk.SkipNode
		},
	}); err != nil {
		log.Errorf("%v\n", err)
		return dirs, err
	}
	var scanner *godirwalk.Scanner
	for _, dir := range allDirs {
		scanner, err = godirwalk.NewScanner(dir)
		if err != nil {
			log.Errorf("cannot lazily scan dir %s: %v\n", dir, err)
			return nil, err
		}
		for scanner.Scan() {
			dirent, err := scanner.Dirent()
			if err != nil {
				log.Warnf("cannot get directory content: %v", err)
				continue
			}
			if dirent.Name() == "Makefile" || dirent.Name() == "statik.go" {
				splitted := strings.Split(dir, "/")
				dirs[splitted[len(splitted)-1]] = dir
			}
		}
	}
	return dirs, err
}
