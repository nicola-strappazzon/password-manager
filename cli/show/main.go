package show

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
	"github.com/nicola-strappazzon/pm/qr"
	"github.com/nicola-strappazzon/pm/term"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

var flagAll bool
var flagCopy bool
var flagQR bool
var flagField string
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "show path/to/encrypted [flags]",
		Short: "Show and decrypt selected data. By default, it shows the password.",
		Example: "  pm show <TAB>\n" +
			"  pm show aws\n" +
			"  pm show aws -p <passphrase>\n" +
			"  pm show aws -p <passphrase> -a\n" +
			"  pm show aws -p <passphrase> -c\n" +
			"  pm show aws -p <passphrase> -f otp -c\n" +
			"  pm show aws -p <passphrase> -f aws.access_key -c",
		Run:               RunCommand,
		ValidArgsFunction: ValidArgs,
	}

	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all decrypted file")
	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().BoolVarP(&flagQR, "qr", "q", false, "Generate a QR code for the decrypted password")
	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Filter by field name...")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	cmd.MarkFlagsMutuallyExclusive("all", "field")
	cmd.MarkFlagsMutuallyExclusive("all", "qr")
	cmd.MarkFlagsMutuallyExclusive("qr", "field")
	cmd.MarkFlagsMutuallyExclusive("qr", "copy")

	cmd.RegisterFlagCompletionFunc("field", FieldCompletion)

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

	if flagAll {
		v = b
	}

	if flagField != "" {
		v = c.Field(flagField)
	}

	if !flagAll && flagField == "" {
		v = c.Password
	}

	if flagField == "otp" {
		v = otp.Get(c.OTP)
	}

	if flagCopy {
		clipboard.Write(v)
		return
	}

	if flagQR {
		qr.Generate(v)
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

func FieldCompletion(cmd *cobra.Command, args []string, toComplete string) (out []string, _ cobra.ShellCompDirective) {
	for _, field := range (&card.Card{}).Fields() {
		if strings.HasPrefix(field, toComplete) {
			out = append(out, field)
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
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
