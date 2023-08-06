package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uzo",
	Short: "Unzip and Open",
	Long: `An CLI Application to unzip .zip files, built with love by Quang Duong.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Hello world.")
	}, 
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
				os.Exit(1)
	}
}

func init() {

}


