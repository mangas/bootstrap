package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/getcouragenow/bootstrap/tool/protofig/config"
	"io/ioutil"
	"log"
	"os"
)

var (
	jsonFile string
	outPath  string
	user     string
)

func main() {
	flag.StringVar(&jsonFile, "f", "./test.json", "json file to decode/encode")
	flag.StringVar(&outPath, "o", "./output", "json output validated against protobuf schema")
	flag.StringVar(&user, "u", "winwisely268", "prefix to output files")
	flag.Parse()

	if jsonFile == "" {
		log.Fatal("Error: you have to supply the jsonfile option")
	}

	if outPath == "" {
		out, err := os.Getwd()
		if err != nil {
			log.Fatal("Cannot get current working dir")
		}
		outPath = out
	}

	if user == "" {
		log.Fatal("Error: user prefix has to be supplied")
	}

	f, err := os.Open(jsonFile)
	if err != nil {
		log.Fatalf("Error: your input json file doesn't exist: %v", err)
	}

	file, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Error: your input can't be read: %v", err)
	}

	var newAppConfig config.DefConfig
	if err := json.Unmarshal(file, &newAppConfig); err != nil {
		log.Fatalf("Error: unable to marshal to json: %v", err)
	}

	if err := createProtojsonOutput(&newAppConfig, outPath, user); err != nil {
		log.Fatalf("Error: failure to create output files: %v", err)
	}
}

func createProtojsonOutput(cfg *config.DefConfig, outpath, user string) (err error) {
	for _, c := range cfg.AppConfig {
		nb, err := c.ConfigCreateJSONMessage()
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(fmt.Sprintf("%s/%s.%s.json", outpath, c.Name, user), nb, 0644); err != nil {
			return err
		}
	}
	return nil
}
