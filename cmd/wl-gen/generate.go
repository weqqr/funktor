package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
	"unicode"
)

//go:embed template.tmpl
var sourceTemplate string

type TemplateInput struct {
	InputPath string
	Copyright string
	Interface Interface
}

func generateSource(input TemplateInput) ([]byte, error) {
	tmpl := template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"comment":    comment,
		"pascalCase": pascalCase,
	})
	tmpl = template.Must(tmpl.Parse(sourceTemplate))

	var writer bytes.Buffer

	err := tmpl.Execute(&writer, input)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func comment(input string) string {
	input = strings.TrimSpace(input)

	var output strings.Builder

	for line := range strings.Lines(input) {
		output.WriteString("// " + strings.TrimSpace(line) + "\n")
	}

	return strings.TrimSpace(output.String())
}

func pascalCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		if len(word) > 0 {
			// Capitalize the first letter of the word
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, "")
}
