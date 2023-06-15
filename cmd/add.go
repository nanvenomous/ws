/*
Copyright © 2023 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"errors"

	"github.com/nanvenomous/ws/system"
	"github.com/spf13/cobra"
)

func runAddCmd(repoUrl string) error {
	var err error
	err = FS.Execute("git", []string{
		"subtree",
		"add",
		"--prefix",
		system.Prefix,
		repoUrl,
		system.Branch,
		"--squash",
	})
	if err != nil {
		return err
	}

	err = FS.Execute("go", []string{
		"work",
		"use",
		system.Prefix,
	})
	if err != nil {
		return err
	}

	return nil
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "clone the repository, add it as a subtree and a workspace module",
	RunE: func(cmd *cobra.Command, args []string) error {
		if system.Prefix == "" {
			return errors.New("must supply flag --prefix")
		}
		if len(args) < 1 {
			return errors.New("Expected 1 argument <repoUrl>")
		}
		return runAddCmd(args[0])
	},
}

func init() {
	system.BranchFlag(addCmd, "b")
	system.PrefixFlag(addCmd, "p")
	rootCmd.AddCommand(addCmd)
}
