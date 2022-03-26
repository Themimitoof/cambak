package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
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
		conf := superchargeConf(cmd, args)

		// Collect all files
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = "Collecting all files..."
		s.Start()

		files := cb_files.CollectFiles(conf)
		allFiles := files.Pictures
		allFiles = append(allFiles, files.RAW...)
		allFiles = append(allFiles, files.Movies...)

		finalMsg := fmt.Sprintf(
			"✅ %d files collected (%d pictures, %d RAWs, %d movies)",
			files.TotalFiles,
			len(files.Pictures),
			len(files.RAW),
			len(files.Movies),
		)
		if conf.Extract.DestinationConflict == cb_config.DEST_CONFLICT_SKIP {
			s.FinalMSG = color.GreenString("%s. %d files skipped.\n", finalMsg, files.SkippedFiles)
		} else {
			s.FinalMSG = color.GreenString(finalMsg + "\n")
		}

		s.Stop()

		// Exit the program if there is no files to copy
		if files.TotalFiles == 0 {
			color.Green("😴 There is no files to copy, please let me sleep.")
			os.Exit(0)
		}

		// Copying files
		color.Yellow("Copying files...")
		bar := progressbar.Default(int64(files.TotalFiles))
		for _, fl := range allFiles {
			if !conf.Extract.DryRunMode {
				destPath, _ := fl.PrepareFileDestinationFolder(conf)
				fl.ExtractFile(conf, destPath)
			}
			bar.Add(1)
		}

		color.Green("✨ All files have been copied!")
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

func superchargeConf(cmd *cobra.Command, args []string) cb_config.Configuration {
	// Manage source path
	if len(args) == 0 {
		color.Red("Please specify the source folder where are located your files.")
		os.Exit(1)
	} else {
		conf.Extract.SourcePath = args[0]
	}

	// Manage destination path
	if len(args) >= 2 {
		conf.Extract.DestinationPath = args[1]
	} else if len(args) < 2 && conf.Extract.DestinationPath == "" {
		color.Red("Please specify the destination folder.")
		os.Exit(1)
	}

	// Manage extraction media types
	allMediasFlag, _ := cmd.Flags().GetBool("all")
	picturesFlag, _ := cmd.Flags().GetBool("pictures")
	rawsFlag, _ := cmd.Flags().GetBool("raws")
	moviesFlag, _ := cmd.Flags().GetBool("movies")

	if allMediasFlag {
		conf.Extract.ExtractPictures = true
		conf.Extract.ExtractRaws = true
		conf.Extract.ExtractMovies = true
	} else if picturesFlag || rawsFlag || moviesFlag {
		conf.Extract.ExtractPictures = picturesFlag
		conf.Extract.ExtractRaws = rawsFlag
		conf.Extract.ExtractMovies = moviesFlag
	}

	// Extraction filters
	// ...

	// Conflict management
	skipFlag, _ := cmd.Flags().GetBool("skip")
	mergeFlag, _ := cmd.Flags().GetBool("merge")

	if skipFlag && mergeFlag {
		color.Red("You can't use --skip and --merge at the same time.")
		os.Exit(1)
	}

	if skipFlag {
		conf.Extract.DestinationConflict = cb_config.DEST_CONFLICT_SKIP
	} else if mergeFlag {
		conf.Extract.DestinationConflict = cb_config.DEST_CONFLICT_MERGE
	}

	// Extraction destination and camera name
	extractFormat, _ := cmd.Flags().GetString("format")
	cameraName, _ := cmd.Flags().GetString("name")

	if extractFormat != "" {
		conf.Extract.DestinationFormat = extractFormat
	}

	if cameraName != "" {
		conf.Extract.CameraName = cameraName
	}

	// Misc flags
	dryRunFlag, _ := cmd.Flags().GetBool("dry-run")
	cleanFlag, _ := cmd.Flags().GetBool("clean")

	conf.Extract.DryRunMode = dryRunFlag

	if cleanFlag {
		conf.Extract.CleanAfterCopy = true
	}

	return conf
}
