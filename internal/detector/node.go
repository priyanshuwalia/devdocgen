package detector

import (
	"os"

	"devdocgen/internal/model"
)

func DetectNode(root string, meta *model.ProjectMetadata) {
	if _, err := os.Stat(root + "/package.json"); err == nil {
		meta.Languages = append(meta.Languages, "JavaScript")
		meta.PackageManager = "npm"
		meta.InstallCommands = append(meta.InstallCommands, "npm install")
		meta.RunCommands = append(meta.RunCommands, "npm start")
	}
}
