package render

import (
	"errors"
	"os"
	"text/template"

	"devdocgen/internal/model"
)

func GenerateReadme(meta *model.ProjectMetadata, outputPath string, overwrite bool, dryRun bool) error {
	tmpl, err := template.ParseFiles("templates/readme.md.tmpl")
	if err != nil {
		return err
	}

	if dryRun {
		return tmpl.Execute(os.Stdout, meta)
	}

	if !overwrite {
		if _, err := os.Stat(outputPath); err == nil {
			return errors.New("output file already exists (use --overwrite to replace)")
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, meta)
}
