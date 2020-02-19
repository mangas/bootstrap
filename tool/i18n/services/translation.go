package services

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/getcouragenow/bootstrap/tool/i18n/utils"
)

// CleanTranslation struct used to clean hugo translations files
type CleanTranslation struct {
	Error *string `json:"error"`
	Fix   *string `json:"fix"`
}

func cleanData(data string, tagsFileDir, tagsFileName string) string {
	cleanedData := data

	cleanTags, _ := getCleanTranslationTags(tagsFileDir, tagsFileName)

	for _, tag := range cleanTags {
		cleanedData = strings.ReplaceAll(cleanedData, *tag.Error, *tag.Fix)
	}
	// for tagIndex, tag := range tagsToRemove {
	// 	cleanedData = strings.ReplaceAll(cleanedData, tag, tagsToReplace[tagIndex])
	// }
	return cleanedData
}

func cleanKey(key string) string {
	return strings.ReplaceAll(key, ".", "_")
}

func getCleanTranslationTags(fileDir, fileName string) ([]*CleanTranslation, error) {

	filePath, err := utils.GetAbsoluteFilePath(fileDir, fileName)

	if err != nil {
		return nil, err
	}

	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	cleanData := []*CleanTranslation{}

	err = json.NewDecoder(f).Decode(&cleanData)
	if err != nil {
		return nil, err
	}

	return cleanData, nil
}

func getTranslateWords(m *linkedhashmap.Map, sep string) string {
	it := m.Iterator()

	out := []string{}
	for it.Next() {
		if !strings.HasPrefix(it.Key().(string), "@") {
			v, ok := it.Value().(string)
			if ok {
				out = append(out, v)
			}
		}
	}
	return strings.Join(out, sep)
}
