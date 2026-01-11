package detector

import (
	"os"

	"devdocgen/internal/model"
)

func DetectDocker(root string, meta *model.ProjectMetadata) {
    if _, err := os.Stat(root + "/Dockerfile"); err == nil {
        meta.HasDocker = true
    }

    if _, err := os.Stat(root + "/docker-compose.yml"); err == nil {
        meta.HasDockerCompose = true
    }
}
	