package cli

import (
	"github.com/nicola-strappazzon/pm/cli/ls"
	"github.com/nicola-strappazzon/pm/cli/completion"

	"github.com/spf13/cobra"
)

func Load() {
	var rootCmd = &cobra.Command{
		Use:  "pm",
		Long: "pm: The standard unix password manager write in go.",
	}

	// rootCmd.CompletionOptions.DisableDefaultCmd = false
	rootCmd.AddCommand(ls.NewCommand())
	rootCmd.AddCommand(completion.NewCommand(rootCmd))
	rootCmd.Execute()
}
