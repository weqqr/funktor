package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("read protocol file: %v", err)
	}

	protocol, err := parseProtocolXML(input)
	if err != nil {
		log.Fatalf("parse protocol XML: %v", err)
	}

	for _, iface := range protocol.Interfaces {
		source, err := generateSource(TemplateInput{
			InputPath: inputPath,
			Copyright: protocol.Copyright,
			Interface: iface,
		})

		if err != nil {
			log.Fatalf("generate source: %v", err)
		}

		path := filepath.Join(
			outputPath,
			protocol.Name,
			iface.Name,
			iface.Name+".go",
		)

		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			log.Fatalf("MkdirAll: %v", err)
		}

		err = os.WriteFile(path, source, 0644)
		if err != nil {
			log.Fatalf("write source: %v", err)
		}
	}
}
