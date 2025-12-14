package otp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/config"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/otp"
	"github.com/nicola-strappazzon/pm/term"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

var flagCopy bool
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "otp path/to/encrypted [flags]",
		Short: "Generate an OTP code and optionally put it on the clipboard.",
		Example: "  pm otp <TAB>\n" +
			"  pm otp aws\n" +
			"  pm otp aws -p <passphrase>\n" +
			"  pm otp aws -p <passphrase> -c\n",
		Run:               RunCommand,
		ValidArgsFunction: ValidArgs,
	}

	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) {
	var v string

	if len(args) == 0 {
		cmd.Help()
		return
	}

	if tree.WalkFrom(GetFirstArg(args)).IsDir {
		tree.WalkFrom(GetFirstArg(args)).Print()
		return
	}

	var b = openpgp.Decrypt(
		term.ReadPassword(flagPassphrase),
		tree.WalkFrom(GetFirstArg(args)).Path,
	)

	var c = card.New(b)

	v = otp.Get(c.OTP)

	if flagCopy {
		clipboard.Write(v)
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), v)
}

func ValidArgs(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
	all, err := ListDirsAndGPG(config.GetDataDirectory())
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
}

func GetFirstArg(in []string) string {
	if len(in) == 1 {
		return in[0]
	}

	return ""
}

func ListDirsAndGPG(basePath string) (list []string, err error) {
	err = filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, err := filepath.Rel(basePath, path)
		if err != nil || rel == "." {
			return nil
		}

		rel = filepath.ToSlash(rel)

		if d.IsDir() {
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
