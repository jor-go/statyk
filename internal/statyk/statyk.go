package statyk

import (
	"fmt"

	"github.com/spf13/cobra"
)

// APP_VERSION Is the current statyk version
const APP_VERSION = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "statyk",
	Short: "statyk is a CLI static website tool",
	Long:  "statyk is a CLI static website tool",
}

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version is the Statyk version",
	Long:    `Version is the Statyk version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(APP_VERSION)
	},
}

func init() {
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(BuildCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(NewCmd)
	rootCmd.AddCommand(ServeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
