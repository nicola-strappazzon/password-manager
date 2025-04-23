package completion

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCommand(rootCmd *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "completion",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

  $ source <(pm completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ pm completion bash > /etc/bash_completion.d/pm
  # macOS:
  $ pm completion bash > /usr/local/etc/bash_completion.d/pm

Zsh:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  $ pm completion zsh > "${fpath[1]}/_pm"
`,
		DisableFlagsInUseLine: true,
		Hidden:                false,
		ValidArgs:             []string{"bash", "zsh"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			switch args[0] {
			case "bash":
				rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				rootCmd.GenZshCompletion(os.Stdout)
			}
		},
	}

	return cmd
}
