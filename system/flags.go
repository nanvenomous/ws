package system

import "github.com/spf13/cobra"

const (
	CHECK_FLAG_NAME  = "check"
	BRANCH_FLAG_NAME = "branch"
	DEFAULT_BRANCH   = "mainline"
)

func CheckFlag(cmd *cobra.Command, flagVar *bool, shorthand string) {
	cmd.Flags().BoolVarP(flagVar, CHECK_FLAG_NAME, shorthand, false, "check the modules to be affected")
}

func BranchFlag(cmd *cobra.Command, flagVar *string, shorthand string) {
	cmd.Flags().StringVarP(flagVar, BRANCH_FLAG_NAME, shorthand, DEFAULT_BRANCH, "specify the branch for the command, default is "+DEFAULT_BRANCH)
}
