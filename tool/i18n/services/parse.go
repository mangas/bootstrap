package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

// TranslatedMaps struct
type TranslatedMaps struct {
	Langs []string
	Maps  []*linkedhashmap.Map
}

// Translate struct
type Translate struct {
	Lang  string
	Words string
}

// ArbAttr struct
type ArbAttr struct {
	Description  string            `json:"description"`
	Type         string            `json:"type"`
	Placeholders map[string]string `json:"placeholders"`
}

// ArbMap convert json to arb
// func ArbMap(data []byte) (*linkedhashmap.Map, error) {
// 	m := linkedhashmap.New()

// 	err := m.FromJSON(data)

// 	if err != nil {
// 		return nil, err
// 	}

// 	it := m.Iterator()
// 	out := linkedhashmap.New()

// 	for it.Next() {

// 		item := ArbAttr{
// 			Description:  "",
// 			Type:         "",
// 			Placeholders: make(map[string]string),
// 		}

// 		out.Put(it.Key(), it.Value())
// 		out.Put("@"+it.Key().(string), item)

// 	}
// 	return out, nil
// }

// JSONMap convert arb to json
func JSONMap(data []byte, full bool) (*linkedhashmap.Map, error) {

	m := linkedhashmap.New()

	err := m.FromJSON(data)
	if err != nil {
		return nil, err
	}
	it := m.Iterator()

	out := linkedhashmap.New()
	for it.Next() {
		if !strings.HasPrefix(it.Key().(string), "@") || full {
			out.Put(it.Key(), it.Value())
		} else {
			if full {
				out.Put(it.Key(), it.Value())
			}
		}
	}
	return out, nil
}

// translate a string from languages to language
func getTemplateWords(m *linkedhashmap.Map, delay time.Duration, tries int, fromLang, sep string, languages []string) ([]Translate, error) {

	words := getTranslateWords(m, sep)
	wordsTranslated := []Translate{}
	for _, lang := range languages {

		t := Translate{}
		out, err := gtranslate.TranslateWithParams(words,
			gtranslate.TranslationParams{
				From:  fromLang,
				To:    lang,
				Tries: tries,
				Delay: delay,
			})
		if err != nil {
			log.Printf("Error to translate from %s to %s\n", fromLang, lang)
		}
		t.Lang = lang
		t.Words = out
		wordsTranslated = append(wordsTranslated, t)
	}
	return wordsTranslated, nil
}

func getTranslatedMaps(sep string, WordsTranslated []Translate, m *linkedhashmap.Map, full bool) (*TranslatedMaps, error) {

	translatedMaps := &TranslatedMaps{}
	for _, tr := range WordsTranslated {
		mapLang := linkedhashmap.New()
		it := m.Iterator()
		index := 0
		words := strings.Split(tr.Words, sep)
		fmt.Println(words)
		fmt.Println(index)
		for it.Next() {
			if !strings.HasPrefix(it.Key().(string), "@") {
				mapLang.Put(it.Key(), strings.TrimSpace(words[index]))
				index++
			} else {
				if full {
					mapLang.Put(it.Key(), it.Value())
				}
			}
		}

		translatedMaps.Maps = append(translatedMaps.Maps, mapLang)
		translatedMaps.Langs = append(translatedMaps.Langs, tr.Lang)

	}
	return translatedMaps, nil
}
