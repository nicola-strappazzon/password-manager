package ls

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/pm/config"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List all encrypted items in tree format.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tree.WalkFrom(GetFirstArg(args)).Print()
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
			all, err := GetAllDirectories(config.GetDataDirectory())
			if err != nil {
				return suggestions, cobra.ShellCompDirectiveNoFileComp
			}

			for _, v := range all {
				if v == toComplete {
					continue
				}

				if strings.HasPrefix(v, toComplete) {
					suggestions = append(suggestions, v)
				}
			}

			return suggestions, cobra.ShellCompDirectiveNoFileComp
		},
	}

	return cmd
}

func GetFirstArg(in []string) string {
	if len(in) == 1 {
		return in[0]
	}

	return ""
}

func GetAllDirectories(root string) (dirs []string, err error) {
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(config.GetDataDirectory(), path)
		if err == nil && rel != "." {
			dirs = append(dirs, rel+"/")
		}
		return nil
	})

	return
}
