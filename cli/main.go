package cli

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/cli/completion"
	"github.com/nicola-strappazzon/password-manager/cli/file"
	"github.com/nicola-strappazzon/password-manager/cli/generate"
	"github.com/nicola-strappazzon/password-manager/cli/ls"
	"github.com/nicola-strappazzon/password-manager/cli/otp"
	"github.com/nicola-strappazzon/password-manager/cli/show"
	"github.com/nicola-strappazzon/password-manager/config"

	"github.com/spf13/cobra"
)

func Load() {
	var rootCmd = &cobra.Command{
		SilenceUsage:      true,
		Use:               "pm",
		Long:              "This is another Unix-style password manager written in Go.",
		PersistentPreRunE: PersistentPreRunE,
		RunE:              RunCommand,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(add.NewCommand())
	rootCmd.AddCommand(completion.NewCommand(rootCmd))
	rootCmd.AddCommand(file.NewCommand())
	rootCmd.AddCommand(generate.NewCommand())
	rootCmd.AddCommand(ls.NewCommand())
	rootCmd.AddCommand(otp.NewCommand())
	rootCmd.AddCommand(show.NewCommand())
	rootCmd.Execute()
}

func PersistentPreRunE(cmd *cobra.Command, args []string) error {
	if config.HasNotPublicKey() {
		return fmt.Errorf("Missing required environment variable: PM_PUBLICKEY")
	}

	if config.HasNotPrivateKey() {
		return fmt.Errorf("Missing required environment variable: PM_PRIVATEKEY")
	}

	return nil
}

func RunCommand(cmd *cobra.Command, args []string) error {
	return ls.NewCommand().RunE(cmd, args)
}
