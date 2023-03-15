/*
Copyright © 2023 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"github.com/nanvenomous/exfs"
	"github.com/nanvenomous/snek"
	"github.com/nanvenomous/ws/configuration"
	"github.com/nanvenomous/ws/exception"
	"github.com/spf13/cobra"
)

const (
	PKG_NAME    = "ws"
	GO_WORKFILE = "go.work"
)

var (
	FS      *exfs.FileSystem
	CFG     *configuration.Configuration
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   PKG_NAME,
	Short: "an opinionated meta repo tool for working with go workspaces and git subtree",
	RunE: func(cmd *cobra.Command, args []string) error {
		snek.RunShellCompletion(cmd)
		return cmd.Help()
	},
}

func Execute() {
	var err error
	err = rootCmd.Execute()
	exception.CheckErr(err)
}

func init() {
	snek.InitRoot(rootCmd, cfgFile, PKG_NAME)
	FS = exfs.NewFileSystem()
	CFG = configuration.NewConfiguration(FS)
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
