package status

import (
	"github.com/nicola-strappazzon/password-manager/internal/git"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the git status of the password store",
		RunE:  RunCommand,
	}
}

func RunCommand(cmd *cobra.Command, args []string) error {
	if !git.IsRepo() {
		cmd.Println("The password store is not a git repository.")
		return nil
	}

	entries := git.StatusEntries()
	commits := git.UnpushedCommits()

	for _, e := range entries {
		cmd.Printf("%s:%-12s %s\n", git.Branch(), e.Label()+":", e.File)
	}

	for _, c := range commits {
		cmd.Printf("%s:%s\n", git.Branch(), c)
	}

	if len(entries) == 0 && len(commits) == 0 {
		cmd.Println("\nNothing to push or commit.")
	}

	return nil
}
