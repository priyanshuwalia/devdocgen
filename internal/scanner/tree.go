package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateTree(root string, indent string) string {
	var sb strings.Builder
	files, _ := os.ReadDir(root)

	for i, file := range files {
		// Ignore hidden files and heavy directories
		if strings.HasPrefix(file.Name(), ".") || file.Name() == "node_modules" {
			continue
		}

		isLast := i == len(files)-1
		prefix := "├── "
		if isLast {
			prefix = "└── "
		}

		sb.WriteString(fmt.Sprintf("%s%s%s\n", indent, prefix, file.Name()))

		if file.IsDir() {
			newIndent := indent
			if isLast {
				newIndent += "    "
			} else {
				newIndent += "│   "
			}
			sb.WriteString(GenerateTree(filepath.Join(root, file.Name()), newIndent))
		}
	}
	return sb.String()
}