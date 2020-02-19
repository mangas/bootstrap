package services

import (
	"fmt"
	"testing"
)

var data = []byte(`
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

func TestTranslate(t *testing.T) {
	fmt.Println(GenerateMultiLanguageFilesFromTemplate("intl.arb", "", "out", ".json", "==", []string{"fr", "es"}, false))
}

func TestGenerateMultiLanguagesFilesFromFiles(t *testing.T) {
	fmt.Println(GenerateMultiLanguagesFilesFromFiles(".", ".", "json", "arb", false))
}
