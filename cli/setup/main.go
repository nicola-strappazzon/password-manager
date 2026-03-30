package setup

import (
	"fmt"
	"os/exec"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "setup",
		Short:             "Configure the application for first use",
		RunE:              RunCommand,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}
}

func RunCommand(cmd *cobra.Command, args []string) error {
	if _, err := exec.LookPath("gpg"); err != nil {
		return fmt.Errorf("gpg is not installed or not found in PATH. Please install GnuPG before using this application.")
	}

	if !config.HasNotRecipient() {
		cmd.Println("The application is already configured.")
		if !term.Confirm("Do you want to continue anyway?") {
			return nil
		}
		cmd.Println("")
	}

	if term.Confirm("Do you already have an OpenPGP key pair?") {
		email := term.ReadLine("What is your e-mail?:")

		cmd.Printf("\nAdd the following lines to your shell profile (~/.zshrc or ~/.bashrc):\n\n")
		cmd.Printf("  export PM_RECIPIENT=\"%s\"\n", email)
		cmd.Println("\nThen reload your shell: source ~/.zshrc")
	}

	return nil
}
