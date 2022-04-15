package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var CambakVersion string
var CambakCommit string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the version of Cambak",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s (commit: %s)\n", color.GreenString("Cambak"), color.GreenString("%s", CambakVersion), CambakCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
