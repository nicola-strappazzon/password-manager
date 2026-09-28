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
		if err := ValidateRecipient(email); err != nil {
			return err
		}

		if err := config.SaveRecipient(email); err != nil {
			return err
		}

		cmd.Printf("Recipient saved in %s\n", config.GetPath(config.GPGIDFile))
	}

	return nil
}

func ValidateRecipient(recipient string) error {
	cmd := exec.Command("gpg", "--list-keys", "--with-colons", fmt.Sprintf("<%s>", recipient))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("the recipient %q was not found in your local GPG keyring", recipient)
	}

	return nil
}
