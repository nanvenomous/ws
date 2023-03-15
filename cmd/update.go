/*
Copyright © 2023 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nanvenomous/ws/system"
	"github.com/spf13/cobra"
)

var (
	syncFlByts       []byte
	syncFlString     string
	updateCheckFlag  bool
	updateBranchFlag string
)

func updateDependencies(deps []string) error {
	var err error
	for _, dep := range deps {
		err = FS.Execute("go", []string{"get", "-u", dep + "@" + updateBranchFlag})
		if err != nil {
			return err
		}
	}
	return nil
}

func fileContains(flToCheck string, toFnd string) (bool, error) {
	var (
		err error
	)
	syncFlByts, err = ioutil.ReadFile(flToCheck)
	if err != nil {
		return false, err
	}
	syncFlString = string(syncFlByts)
	if strings.Contains(syncFlString, toFnd) {
		return true, nil
	}
	return false, nil
}

func runUpdateCmd() error {
	var (
		err    error
		hasDep bool
		deps   []string
	)
	for _, curWf := range CFG.Modules() {
		err = os.Chdir(curWf.Path)
		if err != nil {
			return err
		}
		deps = []string{}
		for _, checkWf := range CFG.Modules() {
			if curWf.Path != checkWf.Path {
				hasDep, err = fileContains("go.mod", checkWf.ModulePath)
				if err != nil {
					return err
				}
				if hasDep {
					deps = append(deps, checkWf.ModulePath)
				}
			}
		}
		if updateCheckFlag {
			fmt.Println(curWf.Path, deps)
		} else {
			fmt.Println("[UPDATING]", curWf.Path)
			err = updateDependencies(deps)
			if err != nil {
				return err
			}
		}
		err = os.Chdir("..")
		if err != nil {
			return err
		}
	}

	return nil
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates all dependencies included in work file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUpdateCmd()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	system.CheckFlag(updateCmd, &updateCheckFlag, "c")
	system.BranchFlag(updateCmd, &updateBranchFlag, "b")
}
