package push

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/internal/git"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "push",
		Short: "Push changes to the remote password store",
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

	if commits := git.UnpushedCommits(); len(commits) == 0 {
		cmd.Println("Nothing to push.")
		return nil
	}

	output, err := git.Push()
	if err != nil {
		return err
	}

	cmd.Print(output)

	return nil
}
