package system

import "github.com/spf13/cobra"

const (
	DEFAULT_BRANCH = "mainline"

	CHECK_FLAG_NAME  = "check"
	BRANCH_FLAG_NAME = "branch"
	PREFIX_FLAG_NAME = "prefix"
)

var (
	Branch string
	Prefix string
)

func CheckFlag(cmd *cobra.Command, flagVar *bool, shorthand string) {
	cmd.Flags().BoolVarP(flagVar, CHECK_FLAG_NAME, shorthand, false, "check the modules to be affected")
}

func BranchFlag(cmd *cobra.Command, shorthand string) {
	cmd.Flags().StringVarP(&Branch, BRANCH_FLAG_NAME, shorthand, DEFAULT_BRANCH, "specify the branch for the command, default is "+DEFAULT_BRANCH)
}

func PrefixFlag(cmd *cobra.Command, shorthand string) {
	cmd.Flags().StringVarP(&Prefix, PREFIX_FLAG_NAME, shorthand, "", "the directory to contain the subtree & workspace")
	cmd.RegisterFlagCompletionFunc(PREFIX_FLAG_NAME, func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	})
}
