package pull

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/internal/git"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "pull",
		Short: "Pull changes from the remote password store",
		RunE:  RunCommand,
	}
}

func RunCommand(cmd *cobra.Command, args []string) error {
	if !git.IsRepo() {
		cmd.Println("The password store is not a git repository.")
		return nil
	}

	if entries := git.StatusEntries(); len(entries) > 0 {
		return fmt.Errorf("uncommitted changes in the password store, commit them first.")
	}

	output, err := git.Pull()
	if err != nil {
		return err
	}

	cmd.Print(output)

	return nil
}
