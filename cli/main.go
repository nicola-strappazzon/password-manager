package cli

import (
	"github.com/nicola-strappazzon/pm/cli/completion"
	"github.com/nicola-strappazzon/pm/cli/ls"
	"github.com/nicola-strappazzon/pm/cli/show"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

func Load() {
	var rootCmd = &cobra.Command{
		Use:  "pm",
		Long: "pm: The standard unix password manager write in go.",
		Run: func(cmd *cobra.Command, args []string) {
			ls.NewCommand().Run(cmd, args)
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return tree.Walk().List(), cobra.ShellCompDirectiveNoFileComp
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = false
	rootCmd.AddCommand(ls.NewCommand())
	rootCmd.AddCommand(completion.NewCommand(rootCmd))
	rootCmd.AddCommand(show.NewCommand())
	rootCmd.Execute()	
}
