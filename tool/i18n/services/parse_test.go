package services_test

import (
	"github.com/getcouragenow/bootstrap/tool/i18n/services"
	"io/ioutil"
	"testing"
)

var (
	data = []byte(`
{
	"title": "Good morning",
	"@title": {
	  "description": "",
	  "type": "",
	  "placeholders": {}
	},
	"title1": "Good",
	"@title1": {
	  "description": "",
	  "type": "",
	  "placeholders": {}
	}
  }
  
`)
)

const (
	success = "\u2713"
	failed  = "\u274c"
)

func TestTranslate(t *testing.T) {
	if err := ioutil.WriteFile("/tmp/data.arb", data, 0644); err != nil {
		t.Fatalf("\t%s\tShould be able to write data file to /tmp: %v", failed, err)
	}
	// fmt.Println(GenerateMultiLanguageFilesFromTemplate("../examples/intl.arb", "", "out", ".json", "==", []string{"fr", "es"}, false))
	t.Run("Test Multi Languages Files From Template", testGenerateMultiLanguageFilesFromTemplate)
}

func testGenerateMultiLanguageFilesFromTemplate(t *testing.T) {
	t.Log("Test multi language files from template")
	{
		err := services.GenerateMultiLanguageFilesFromTemplate(
			"/tmp/data.arb",
			"/tmp",
			"out",
			".json",
			"==",
			[]string{"fr", "es"},
			false,
		)
		if err != nil {
			t.Fatalf(
				"\t%s\tShould be able to write translated output files to /tmp, got: %v",
				failed,
				err,
			)
		}
		t.Logf("\t%s\tShould be able to write translated output files to /tmp",
			success)
	}
}

