package version

import (
	"github.com/spf13/cobra"
)

var VERSION string

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Print version number",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(VERSION)
		},
	}

	return cmd
}
