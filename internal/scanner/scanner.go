package scanner

import (
	"io/fs"
	"path/filepath"
)

type RepoSummary struct {
	Files       []string
	Directories []string
}

func Scan(root string) (*RepoSummary, error) {
	summary := &RepoSummary{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			summary.Directories = append(summary.Directories, path)
		} else {
			summary.Files = append(summary.Files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return summary, nil
}
