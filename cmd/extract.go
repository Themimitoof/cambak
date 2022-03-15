package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	cb_config "github.com/themimitoof/cambak/config"
	cb_files "github.com/themimitoof/cambak/files"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Copy files from a source media to a local/remote destination",
	Long: `The cambak extrator will copy/extract files from a source media (eg:
SD card, MTP drive, local/remote folder) to a local or remote destination folder.

By default, the folder destination structure is the following:

<destination folder>
└── <YEAR>
    └── <MONTH>-<DAY>
        └── <CAMERA_NAME>
            ├── Pictures
            ├── RAW
            └── Movies

You can change the destination format by using the '--format' flag or change the value
in the configuration file.

For more information, please consult: https://github.com/themimitoof/cambak.`,
	Aliases: []string{"copy", "cp"},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plop")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// File extraction selectors
	extractCmd.Flags().BoolP("all", "A", false, "Import all medias files type")
	extractCmd.Flags().BoolP("pictures", "P", false, "Import pictures files")
	extractCmd.Flags().BoolP("raws", "R", false, "Import RAWs files")
	extractCmd.Flags().BoolP("movies", "M", false, "Import movies files")
	// extractCmd.Flags().StringP("date", "d", "", "Specify a date or a range to extract (not yet implemented)")
	// extractCmd.Flags().StringP("rate", "r", "", "Specify a rate or a rating range to extract (not yet implemented)")

	// Extraction collision management
	extractCmd.Flags().BoolP("skip", "s", false, "Skip the source file if it already exists in the destination folder")
	extractCmd.Flags().BoolP("merge", "m", false, "Merge the source file if it already exists in the destination folder")

	extractCmd.Flags().StringP("format", "f", "", "Structure format in the destination folder.")
	extractCmd.Flags().StringP("name", "n", "", "Name of the camera")

	// Misc
	extractCmd.Flags().Bool("dry-run", false, "Only log what the extractor will do if this flag was not set")
	extractCmd.Flags().BoolP("clean", "c", false, "Delete source file after been copied")
	// extractCmd.Flags().BoolP("auto-source", "a", false, "Auto guess the source file (not yet implemented)")
	// extractCmd.Flags().BoolP("format-source", "z", false, "Format the source after the extraction done (not yet implemented)")
}
