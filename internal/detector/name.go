package detector

import (
	"path/filepath"

	"devdocgen/internal/model"
)

func DetectName(root string, meta *model.ProjectMetadata) {
    absPath, err := filepath.Abs(root)
    if err != nil {
        meta.Name = root
        return
    }

    meta.Name = filepath.Base(absPath)
}
