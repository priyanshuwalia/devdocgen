package detector

import (
	"os"
	"path/filepath"
	"strings"

	"devdocgen/internal/model"
)
func readFile(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        return ""
    }
    return string(content)
}
func DetectGo(root string, meta *model.ProjectMetadata) {
    modPath := filepath.Join(root, "go.mod")
    if _, err := os.Stat(modPath); err == nil {
        meta.Languages = append(meta.Languages, "Go")
        
        content := readFile(modPath)
        // Detect Frameworks
        if strings.Contains(content, "github.com/spf13/cobra") {
            meta.Languages = append(meta.Languages, "Cobra (CLI)")
        }
        if strings.Contains(content, "github.com/gin-gonic/gin") {
            meta.Languages = append(meta.Languages, "Gin (Web Framework)")
        }
    }
}