package show

import (
	"fmt"
	// "github.com/nicola-strappazzon/pm/tree"
	"github.com/nicola-strappazzon/pm/openpgp"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show and decrypt selected data.",
		Run: func(cmd *cobra.Command, args []string) {
			// nodes := tree.WalkFrom("")
			// fmt.Println(nodes.Exist(FirstOrEmpty(args)))

			if !openpgp.IsGPGInstalled() {
				fmt.Println("GPG is NOT Installed.")
				return
			}

			content, err := openpgp.Decrypt("/Users/nicola/.password-store/thn/vpn.gpg")

			fmt.Println(err)
			fmt.Println(content)
		},
	}

	return cmd
}

func FirstOrEmpty(in []string) string {
	if len(in) > 0 {
		return in[0]
	}
	return ""
}
