package cli

import (
	"github.com/nicola-strappazzon/pm/cli/completion"
	"github.com/nicola-strappazzon/pm/cli/ls"
	"github.com/nicola-strappazzon/pm/cli/otp"
	"github.com/nicola-strappazzon/pm/cli/show"

	"github.com/spf13/cobra"
)

func Load() {
	var rootCmd = &cobra.Command{
		Use:  "pm",
		Long: "This is another Unix-style password manager written in Go.",
		Run: func(cmd *cobra.Command, args []string) {
			ls.NewCommand().Run(cmd, args)
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(ls.NewCommand())
	rootCmd.AddCommand(completion.NewCommand(rootCmd))
	rootCmd.AddCommand(show.NewCommand())
	rootCmd.AddCommand(otp.NewCommand())
	rootCmd.Execute()
}
