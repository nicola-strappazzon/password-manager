package show

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

var passphrase string
var clip bool

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show and decrypt selected data.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			if tree.WalkFrom(GetFirstArg(args)).IsDir {
				tree.WalkFrom(GetFirstArg(args)).Print()
				return
			}

			Run(cmd, tree.WalkFrom(GetFirstArg(args)).Path)
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return tree.Walk().List(), cobra.ShellCompDirectiveNoFileComp
		},
	}

	cmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")
	cmd.Flags().BoolVarP(&clip, "clip", "c", false, "Copy decrypted data to clipboard")

	return cmd
}

func Run(cmd *cobra.Command, path string) {
	content := openpgp.Decrypt(
		passphrase,
		path,
	)

	if clip {
		clipboard.Write(content)
	} else {
		fmt.Fprintln(cmd.OutOrStdout(), content)
	}

	return
}

func GetFirstArg(in []string) string {
	if len(in) == 1 {
		return in[0]
	}

	return ""
}
