package detector

import (
	"os"
	"path/filepath"
	"strings"
)

func DetectLicense(root string) string {
	candidates := []string{
		"LICENSE",
		"LICENSE.md",
		"LICENSE.txt",
	}

	for _, name := range candidates {
		path := filepath.Join(root, name)

		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		content := strings.ToLower(string(data[:min(len(data), 2000)]))

		switch {
		case strings.Contains(content, "mit license"):
			return "MIT"
		case strings.Contains(content, "apache license"):
			return "Apache-2.0"
		case strings.Contains(content, "gnu general public license"):
			return "GPL"
		case strings.Contains(content, "bsd license"):
			return "BSD"
		}
	}

	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
