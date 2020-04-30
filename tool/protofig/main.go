package main

import (
	"encoding/json"
	"github.com/getcouragenow/boostrap/tool/protofig/config"
	"io/ioutil"
	"log"
)

type intermediateType struct {
	componentName string                   `json:"componentName"`
	config        []map[string]interface{} `json:"config"`
}

/*
The first part:
1. given a file, likely json with a key and value, validates them against baseproto's Config and ConfigVal
2. create env-${USER}.json from that protobuf.
*/

func newIty(file string) (ity *intermediateType, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(f, ity); err != nil {
		return nil, err
	}
	return ity, nil
}

func (ity *intermediateType) validateComponentName() {
	validComponentName := config.HasMessageName(ity.componentName)
	if !validComponentName {
		log.Fatal("Component is invalid")
	}
	comptype := config.CreateMessage(ity.componentName)
	log.Println(comptype)
}

func main() {
	ity, err := newIty("test.json")
	if err != nil {
		log.Fatal(err)
	}
	ity.validateComponentName()
}
