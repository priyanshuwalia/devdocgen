package detector

import (
	"os"

	"devdocgen/internal/model"
)

func DetectEnv(root string, meta *model.ProjectMetadata) {
    if _, err := os.Stat(root + "/.env"); err == nil {
        meta.HasEnvFile = true
    }

    if _, err := os.Stat(root + "/.env.example"); err == nil {
        meta.HasEnvFile = true
        meta.EnvExampleFile = ".env.example"
    }
}
