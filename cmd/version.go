package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const CAMBAK_VERSION string = "0.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the version of Cambak",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s\n", color.GreenString("Cambak"), color.GreenString("%s", CAMBAK_VERSION))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
