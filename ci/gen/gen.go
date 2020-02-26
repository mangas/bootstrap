package gen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/getcouragenow/bootstrap/ci/gen/template"
	"github.com/magefile/mage/mg"
)

type libInfo struct {
	text         string
	defaultValue string
	input        string
}

var (
	mageFileName = "gen_mage.go"
	infos        = []libInfo{
		{"LibName", "main", ""},
		{"Lib", "github.com/getcouragenow/", "github.com/getcouragenow/"},
		{"Branch", "master", "master"},
	}

	libName = ""
	lib     = ""
	branch  = ""
)

// Gen namespace
type Gen mg.Namespace

func getMageInfos() {

	for _, info := range infos {
		reader := bufio.NewReader(os.Stdin)
		d := fmt.Sprintf("(default: %s): ", info.defaultValue)
		fmt.Print(info.text, d)
		info.input, _ = reader.ReadString('\n')

		if info.text == "LibName" {
			libName = strings.TrimRight(info.input, "\n")
		} else if info.text == "Lib" {
			lib = strings.TrimRight(info.input, "\n")
		} else if info.text == "Branch" {
			branch = strings.TrimRight(info.input, "\n")
		}
	}
	fmt.Println(libName)
	fmt.Println(lib)
	fmt.Println(branch)
}

// MageFile generate mage file with basic setup.
func (Gen) MageFile() error {

	getMageInfos()
	file, err := os.Create(mageFileName)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(template.GetTemplate(libName, lib, branch)))
	return nil
}
