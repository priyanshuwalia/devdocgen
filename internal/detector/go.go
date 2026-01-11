package detector

import (
	"os"

	"devdocgen/internal/model"
)

func DetectGo(root string, meta *model.ProjectMetadata) {
	if _, err := os.Stat(root + "/go.mod"); err == nil {
		meta.Languages = append(meta.Languages, "Go")
		meta.InstallCommands = append(meta.InstallCommands, "go mod tidy")
		meta.RunCommands = append(meta.RunCommands, "go run .")
	}
}
