/*
Copyright © 2023 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all modules from go.work",
	Run: func(cmd *cobra.Command, args []string) {
		for _, wf := range CFG.Modules() {
			fmt.Println(wf.Path, wf.ModulePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
