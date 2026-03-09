package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
	"github.com/ProtonMail/gopenpgp/v3/profile"
	"github.com/nicola-strappazzon/password-manager/config"
	"github.com/nicola-strappazzon/password-manager/file"
	"github.com/nicola-strappazzon/password-manager/term"

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
	if !config.HasNotPublicKey() && !config.HasNotPrivateKey() {
		cmd.Println("The application is already configured.")
		if !term.Confirm("Do you want to continue anyway?") {
			return nil
		}
		cmd.Println("")
	}

	if term.Confirm("Do you already have an OpenPGP key pair?") {
		return importKeys(cmd)
	}

	return generateKeys(cmd)
}

func expandPath(p string) string {
	if strings.HasPrefix(p, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return p
		}
		return filepath.Join(home, p[2:])
	}
	return p
}

func importKeys(cmd *cobra.Command) error {
	pubPath := expandPath(term.ReadLine("Path to public key:"))
	privPath := expandPath(term.ReadLine("Path to private key:"))

	if _, err := crypto.NewKeyFromArmored(file.ReadInString(pubPath)); err != nil {
		return fmt.Errorf("Invalid public key: %w", err)
	}

	if _, err := crypto.NewKeyFromArmored(file.ReadInString(privPath)); err != nil {
		return fmt.Errorf("Invalid private key: %w", err)
	}

	printExportInstructions(cmd, pubPath, privPath)
	return nil
}

func generateKeys(cmd *cobra.Command) error {
	name := term.ReadLine("Name:")
	email := term.ReadLine("Email:")

	passphrase := term.ReadPassword("Passphrase: ", "")
	confirm := term.ReadPassword("Confirm passphrase: ", "")

	if passphrase != confirm {
		return fmt.Errorf("Passphrases do not match.")
	}

	cmd.Println("\nGenerating key pair...")

	pgp := crypto.PGPWithProfile(profile.Default())
	key, err := pgp.KeyGeneration().AddUserId(name, email).New().GenerateKey()
	if err != nil {
		return err
	}
	defer key.ClearPrivateParams()

	lockedKey, err := crypto.PGP().LockKey(key, []byte(passphrase))
	if err != nil {
		return err
	}

	pubKey, err := key.ToPublic()
	if err != nil {
		return err
	}

	pubArmored, err := pubKey.Armor()
	if err != nil {
		return err
	}

	privArmored, err := lockedKey.Armor()
	if err != nil {
		return err
	}

	pubPath := config.GetPath("public.asc")
	privPath := config.GetPath("private.asc")

	if file.Exists(pubPath) || file.Exists(privPath) {
		cmd.Printf("Key files already exist in %s.\n", config.GetPath(""))
		if !term.Confirm("Overwrite?") {
			return nil
		}
	}

	if err := os.MkdirAll(filepath.Dir(pubPath), 0700); err != nil {
		return err
	}

	if err := os.WriteFile(pubPath, []byte(pubArmored), 0600); err != nil {
		return err
	}

	if err := os.WriteFile(privPath, []byte(privArmored), 0600); err != nil {
		return err
	}

	cmd.Printf("  Public key saved to %s\n", pubPath)
	cmd.Printf("  Private key saved to %s\n", privPath)

	printExportInstructions(cmd, pubPath, privPath)
	return nil
}

func printExportInstructions(cmd *cobra.Command, pubPath, privPath string) {
	cmd.Printf("\nAdd the following lines to your shell profile (~/.zshrc or ~/.bashrc):\n\n")
	cmd.Printf("  export PM_PUBLICKEY=\"%s\"\n", pubPath)
	cmd.Printf("  export PM_PRIVATEKEY=\"%s\"\n", privPath)
	cmd.Println("\nThen reload your shell: source ~/.zshrc")
}
