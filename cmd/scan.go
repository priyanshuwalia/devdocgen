package cmd

import (
	"fmt"
	"path/filepath"

	"devdocgen/internal/detector"
	"devdocgen/internal/render"
	"devdocgen/internal/scanner"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	outputPath string
	overwrite  bool
	dryRun     bool
)

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a repository and generate a README",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
    path := args[0]

  
    _, err := scanner.Scan(path)
    if err != nil {
        return err
    }

   
    meta := detector.DetectProject(path)
    meta.Structure = scanner.GenerateTree(path, "")


    if meta.Name == "" || meta.Name == "." {
        var name string
        prompt := &survey.Input{
            Message: "What is the name of this project?",
            Default: filepath.Base(path),
        }
     
        if err := survey.AskOne(prompt, &name); err != nil {
            return err
        }
        meta.Name = name
    }


    err = render.GenerateReadme(meta, outputPath, overwrite, dryRun)
    if err != nil {
        return err
    }

    if !dryRun {
        fmt.Println("README generated at:", outputPath)
    }

    return nil
},
}

func init() {
	scanCmd.Flags().StringVarP(
		&outputPath,
		"output",
		"o",
		"README.generated.md",
		"Output README file",
	)
	scanCmd.Flags().BoolVar(
		&overwrite,
		"overwrite",
		false,
		"Overwrite existing README",
	)
	scanCmd.Flags().BoolVar(
		&dryRun,
		"dry-run",
		false,
		"Print README to stdout without writing file",
	)

	rootCmd.AddCommand(scanCmd)
}
