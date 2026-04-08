package git

import (
	"github.com/nicola-strappazzon/password-manager/cli/git/pull"
	"github.com/nicola-strappazzon/password-manager/cli/git/push"
	"github.com/nicola-strappazzon/password-manager/cli/git/status"
	"github.com/nicola-strappazzon/password-manager/internal/git"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Manage the git repository of the password store",
	}

	cmd.AddCommand(pull.NewCommand())
	cmd.AddCommand(push.NewCommand())
	cmd.AddCommand(status.NewCommand())

	return cmd
}

func Warnings() []string {
	var warnings []string

	if msg := git.UncommittedChangesWarning(); msg != "" {
		warnings = append(warnings, msg)
	}

	if msg := git.UnpushedCommitsWarning(); msg != "" {
		warnings = append(warnings, msg)
	}

	return warnings
}
