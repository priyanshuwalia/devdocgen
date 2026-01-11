package render

import (
	"os"
	"text/template"

	"devdocgen/internal/model"
)

func GenerateReadme(meta *model.ProjectMetadata, outputPath string) error {
	tmpl, err := template.ParseFiles("templates/readme.md.tmpl")
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, meta)
}