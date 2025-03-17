package services

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "In this version, you can acess multiple funcionalities, like register NFEs and control it status!!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NFe controller - v1.0 - HEAD")
	},
}
