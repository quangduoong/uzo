/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/quangduoong/uzo/helper"
	"github.com/spf13/cobra"
)

var (
	srcPath         string
	destPath        string
	isOpenAfterward bool
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "Unzip tool",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if srcPath == "" {
			if len(args) == 0 {
				fmt.Print("User did not provide a file path.")
				return
			}
			srcPath = args[0]
		}

		if !helper.IsZipFile(srcPath) {
			fmt.Printf("File \"%v\" is not a zip file.", srcPath)
			return
		}

		if !helper.IsFileExists(srcPath) {
			fmt.Printf("File \"%v\" does not exist.", srcPath)
			return
		}

		if destPath == "" {
			destPath = strings.Split(srcPath, ".zip")[0]
		} else if helper.IsDirExists(destPath) && !filepath.IsAbs(destPath) {
			fmt.Printf("Directory \"%v\" already exists.", destPath)
			return
		}

		if err := helper.Unzip(srcPath, destPath); err != nil {
			fmt.Println(err)
			return
		}

		if isOpenAfterward {
			if err := helper.OpenInExplorer(destPath); err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	unzipCmd.Flags().StringVarP(&srcPath, "src", "s", "", "File's source path")
	unzipCmd.Flags().StringVarP(&destPath, "dest", "d", "", "File's destination path")
	unzipCmd.Flags().BoolVarP(&isOpenAfterward, "open", "o", false, "Open directory after unzipping completed.")
	rootCmd.AddCommand(unzipCmd)
}
