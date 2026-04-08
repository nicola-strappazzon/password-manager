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

	if branch := git.Branch(); branch != "" {
		cmd.Printf("Branch: %s\n", branch)
	}

	entries := git.StatusEntries()
	if len(entries) > 0 {
		cmd.Println("\nUncommitted changes:")
		for _, e := range entries {
			cmd.Printf("  %-12s %s\n", e.Label()+":", e.File)
		}
	}

	commits := git.UnpushedCommits()
	if len(commits) > 0 {
		cmd.Println("\nUnpushed commits:")
		for _, c := range commits {
			cmd.Printf("  %s\n", c)
		}
	}

	if len(entries) == 0 && len(commits) == 0 {
		cmd.Println("\nNothing to push or commit.")
	}

	return nil
}
