package internal

import (
	"github.com/spf13/cobra"

	"rfm-nfe-cli/internal/services"
)

var rootCmd = &cobra.Command{
	Use:   "nfe-cli",
	Short: "Hello, i'm nfe controller",
}

func Execute() error {
	return rootCmd.Execute()
}

func Init() {
	rootCmd.AddCommand(services.VersionCmd)
}
