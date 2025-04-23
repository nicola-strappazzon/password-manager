package ls

import (
	"github.com/nicola-strappazzon/pm/suggestions"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List passwords.",
		Run: func(cmd *cobra.Command, args []string) {
			nodes := tree.Build(FirstOrEmpty(args))
			nodes.Print()
		},
		Args: cobra.ArbitraryArgs,
	}

	cmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		options := suggestions.Suggestions(toComplete)

		var completions []string
		for _, opt := range options {
			if len(toComplete) == 0 || startsWith(opt, toComplete) {
				completions = append(completions, opt)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	return cmd
}

func FirstOrEmpty(in []string) string {
	if len(in) > 0 {
		return in[0]
	}
	return ""
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
