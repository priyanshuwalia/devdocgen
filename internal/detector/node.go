package detector

import (
	"os"
	"path/filepath"
	"strings"

	"devdocgen/internal/model"
)

func DetectNode(root string, meta *model.ProjectMetadata) {
    pkgPath := filepath.Join(root, "package.json")
    if _, err := os.Stat(pkgPath); err == nil {
        content := readFile(pkgPath)
        meta.Languages = append(meta.Languages, "Node.js")

        // Detect Frontend Frameworks
        if strings.Contains(content, "\"next\"") {
            meta.Languages = append(meta.Languages, "Next.js")
        } else if strings.Contains(content, "\"react\"") {
            meta.Languages = append(meta.Languages, "React")
        }

        // Detect Web3/Blockchain
        if strings.Contains(content, "\"ethers\"") || strings.Contains(content, "\"hardhat\"") {
            meta.Languages = append(meta.Languages, "Web3 (Ethers/Hardhat)")
        }
        
        // Detect Styling
        if strings.Contains(content, "\"tailwindcss\"") {
            meta.Languages = append(meta.Languages, "Tailwind CSS")
        }
    }
}