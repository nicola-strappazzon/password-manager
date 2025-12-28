package completion

import (
	"strings"

	"github.com/nicola-strappazzon/password-manager/card"
	"github.com/nicola-strappazzon/password-manager/explorer"

	"github.com/spf13/cobra"
)

func SuggestDirectories(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	items, err := explorer.Directories()

	if err != nil {
		return suggestions, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}

	return Suggestions(items, toComplete)
}

func SuggestDirectoriesAndFiles(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	items, err := explorer.DirectoriesAndFiles()

	if err != nil {
		return suggestions, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}

	return Suggestions(items, toComplete)
}

func Suggestions(items []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	for _, v := range items {
		if v == toComplete {
			continue
		}

		if strings.HasPrefix(v, toComplete) {
			suggestions = append(suggestions, v)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func SuggestFields(cmd *cobra.Command, args []string, toComplete string) (out []string, _ cobra.ShellCompDirective) {
	for _, field := range (&card.Card{}).Fields() {
		if strings.HasPrefix(field, toComplete) {
			out = append(out, field)
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
