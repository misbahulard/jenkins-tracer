package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version of application
// TODO don't forget to update the version if you are doing any changes on it,
//      we are using a semantic versioning on this project, https://semver.org/
const Version string = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Jenkins Tracer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
