package main

import (
	"bytes"
	"flag"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/osutil"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
	"text/template"
)

var (
	toolsPath  string
	statikPath string
	outputPath string
	gotpl      = `package cmd
	
var (
	statikNamespaces = []string{
		{{ range $k, $v := .Statiks }}"{{ $k }}",
		{{ end }}
	}
	availableTools = []string{
		{{ range $k, $v := .Tools }}"{{ $k }}",
		{{ end }}
	}
	statikPathMap = map[string]string{
		{{ range $k, $v := .Statiks }}"{{ $k }}": "{{ $v }}",
		{{ end }}
	}
	toolPathMap = map[string]string{
		{{ range $k, $v := .Statiks }}"{{ $k }}": "{{ $v }}",
		{{ end }}
	}
)	
	`
)

type StatikTools struct {
	Statiks map[string]string
	Tools   map[string]string
}

func main() {
	flag.StringVar(&statikPath, "s", "${GOPATH}/src/github.com/getcouragenow/bootstrap/statiks",
		`scans statiks path`)
	flag.StringVar(&toolsPath, "t", "${GOPATH}/src/github.com/getcouragenow/bootstrap/tools",
		`scans tools path`)
	flag.StringVar(&outputPath, "o", "${GOPATH}/src/github.com/getcouragenow/bootstrap/sdk/cmd/paths.go",
		`output go file`)
	flag.Parse()

	if statikPath == "" {
		log.Fatal("error: -s flag cannot be empty")
	}
	if toolsPath == "" {
		log.Fatal("error: -t flag cannot be empty")
	}
	if outputPath == "" {
		log.Fatal("error: -o flag cannot be empty")
	}

	statikDirs, err := osutil.WalkPath(statikPath)
	if err != nil {
		log.Fatalf("error scanning %s dir: %v\n", statikPath, err)
	}
	toolsDir, err := osutil.WalkPath(toolsPath)
	if err != nil {
		log.Fatalf("error scanning %s dir: %v\n", toolsPath, err)
	}
	s := &StatikTools{statikDirs, toolsDir}
	b := bytes.NewBuffer(nil)
	fmap := template.FuncMap{
		"getLast": getLastString,
	}
	t, err := template.New("paths").Funcs(fmap).Parse(gotpl)
	if err != nil {
		log.Fatalf("error reading default template: %v\n", err)
	}
	if err := t.Execute(b, s); err != nil {
		log.Fatalf("error executing golang template: %v\n", err)
	}
	if err := ioutil.WriteFile(outputPath, b.Bytes(), 0644); err != nil {
		log.Fatalf("error writing file to %s: %v\n", outputPath, err)
	}
}

func getLastString(s string) string {
	paths := strings.Split(s, "/")
	return paths[len(paths)-1]
}
