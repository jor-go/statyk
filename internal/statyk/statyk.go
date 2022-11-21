package statyk

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "statyk",
	Short: "statyk is a CLI static website tool",
	Long:  "statyk is a CLI static website tool",
}

func init() {
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
