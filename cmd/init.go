package cmd

import (
	"github.com/lanlin/notion-site/pkg"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init the config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.DefaultConfigInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
