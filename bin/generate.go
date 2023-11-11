package main

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	OutputFile         = "./DerivApi.yml"
	SchemaPath         = "./deriv-developers-portal/config/v3/"
	RequestExampleFile = "example.json"
	FormatVersion      = "1"
)

type Config struct {
	Version string              `yaml:"version"`
	Domains []string            `yaml:"domains"`
	Macro   map[string][]string `yaml:"macro"`
}

func main() {
	config := Config{Version: FormatVersion}
	config.Domains = []string{"derivws.com", "binaryws.com", "deriv.dev"}
	config.Macro = make(map[string][]string)

	files, err := os.ReadDir(SchemaPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()

		rawExample, err := os.ReadFile(SchemaPath + name + "/" + RequestExampleFile)
		if err != nil {
			log.Fatal(err)
		}

		example := strings.TrimSpace(string(rawExample))

		config.Macro[name] = []string{"edit " + example}
	}

	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	os.WriteFile(OutputFile, yamlData, 0644)
}
