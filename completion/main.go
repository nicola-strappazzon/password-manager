package completion

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/config"

	"github.com/spf13/cobra"
)

func SuggestDirectories(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	items, err := Directories()

	if err != nil {
		return suggestions, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}

	return Suggestions(items, toComplete)
}

func Directories() (dirs []string, err error) {
	basePath := config.GetDataDirectory()

	err = filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			return nil
		}

		rel, _ := filepath.Rel(basePath, path)
		if rel == "." {
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") {
			return filepath.SkipDir
		}

		dirs = append(dirs, rel+"/")

		return nil
	})

	return
}

func SuggestDirectoriesAndFiles(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	items, err := DirectoriesAndFiles()

	if err != nil {
		return suggestions, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}

	return Suggestions(items, toComplete)
}

func DirectoriesAndFiles() (list []string, err error) {
	basePath := config.GetDataDirectory()

	err = filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(basePath, path)
		if rel == "." {
			return nil
		}

		rel = filepath.ToSlash(rel)

		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}

			if !strings.HasSuffix(rel, "/") {
				rel += "/"
			}

			list = append(list, rel)
			return nil
		}

		if strings.HasSuffix(d.Name(), ".gpg") {
			noExt := strings.TrimSuffix(rel, ".gpg")
			list = append(list, noExt)
		}

		return nil
	})

	return
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
