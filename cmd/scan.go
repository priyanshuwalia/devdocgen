package cmd

import (
	"fmt"

	"devdocgen/internal/detector"
	"devdocgen/internal/render"
	"devdocgen/internal/scanner"

	"github.com/spf13/cobra"
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

		err = render.GenerateReadme(meta, "README.generated.md")
		if err != nil {
			return err
		}

		fmt.Println("README.generated.md created successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
