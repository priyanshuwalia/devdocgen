package cmd

import (
	"fmt"

	"devdocgen/internal/detector"
	"devdocgen/internal/render"
	"devdocgen/internal/scanner"

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
