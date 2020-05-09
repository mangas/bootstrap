package statics

/*
RootBoilerplate satisfies AssetEmbedder interface
This one is namespaced for root of the boilerplate directory
see boilerplate directory in bootstrap to see the content.
 */

import (
	"fmt"
	_ "github.com/getcouragenow/bootstrap/statik"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const boilerplateNS = "rootBoilerplate"

type RootBoilerplate struct {
	fsys http.FileSystem // the rakyll fs
}

// NewRootBoilerplate will return RootBoilerplate
func NewRootBoilerplate() (*RootBoilerplate, error){
	bfs, err := fs.NewWithNamespace(boilerplateNS)
	if err != nil {
		return nil, err
	}
	return &RootBoilerplate{
		fsys:    bfs,
	}, nil
}

func (r *RootBoilerplate) GetFs() http.FileSystem { return r.fsys }
func (r *RootBoilerplate) WriteAllFiles(outputPath string) error {
	if err := checkAndMakeDir(outputPath); err != nil {
		return err
	}
	if err := fs.Walk(r.fsys, "/boilerplates", func(filePath string, fileInfo os.FileInfo, err error) error {
		newPath := path.Join("/boilerplate", filePath)
		if fileInfo.IsDir() {
			if err := checkAndMakeDir(newPath); err != nil {
				return fmt.Errorf("creating directory %q: %w", newPath, err)
			}
		} else {
			file, err := r.fsys.Open(filePath)
			if err != nil {
				return fmt.Errorf("opening %q in embedded filesystem: %w", filePath, err)
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return fmt.Errorf("reading %q in embedded filesystem: %w", filePath, err)
			}

			if err := ioutil.WriteFile(newPath, buf, 0664); err != nil {
				return fmt.Errorf("writing %q to %q: %w", filePath, newPath, err)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	s, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}

	log.Printf("Successfully exported boilerplates to %s", s)
	return nil
}

func (r *RootBoilerplate) ReadSingleFile(name string) ([]byte, error) {
	f, err := r.fsys.Open(fmt.Sprintf("/%s", name))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
