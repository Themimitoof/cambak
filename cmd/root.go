package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/themimitoof/cambak/config"
)

var confPath string
var defaultConfPath string
var conf config.Configuration

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cambak",
	Short: "A simple but powerful tool for derushing cameras",
	Long: `Cambak is a simple but powerful too for derushing cameras.

The program use a configuration file located in '$HOME/.config/cambak.yaml'.
During the first execution, a default configuration file will be created. You
can override it by an another configuration file by using the --config flag.

For more information, please consult: https://github.com/themimitoof/cambak.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		userConfPath, _ := cmd.Flags().GetString("config")

		// If no custom configuration file have been specified and if this one doesn't
		// exists, create it.
		if userConfPath == defaultConfPath {
			_, err := os.Stat(userConfPath)

			if err != nil {
				fmt.Println("Generating the default configuration file...")
				err = config.NewConfigurationFile(userConfPath)

				if err != nil {
					fmt.Printf("Unable to generate the default configuration file. Err: %s\n", err)
					os.Exit(255)
				}
			}
		}

		// Load the configuration
		var err error
		conf, err = config.OpenConfigurationFile(userConfPath)

		if err != nil {
			fmt.Printf("Unable to open or read the configuration file. Err: %s\n", err)
			os.Exit(2)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Get the $HOME value to build the value of the `defaultConfPath` var
	environ := os.Environ()
	homeFolder := "/tmp" // Just in case the $HOME env doesn't exists (for super sandboxed/weird envs)

	for _, v := range environ {
		splittedVal := strings.Split(v, "=")
		key := splittedVal[0]
		val := splittedVal[1]

		if key == "HOME" {
			homeFolder = val
			break
		}
	}
	defaultConfPath = homeFolder + "/.config/cambak.yaml"

	// Specify global flags
	rootCmd.PersistentFlags().StringVar(&confPath, "config", defaultConfPath, "Path of the configuration file")
}
