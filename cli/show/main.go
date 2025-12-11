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
var flagClip bool
var flagQR bool
var flagField string
var flagPassphrase string

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show path/to/encrypted [flags]",
		Short: "Show and decrypt selected data. By default, it shows the password.",
		Example: "  pm show <TAB>\n" +
			"  pm show wifi/theforce -p <passphrase>\n" +
			"  pm show wifi/theforce -p <passphrase> -a\n" +
			"  pm show wifi/theforce -p <passphrase> -c\n" +
			"  pm show com/aws -p <passphrase> -f otp -c\n" +
			"  pm show com/aws -p <passphrase> -f aws.access_key -c",
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
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (suggestions []string, _ cobra.ShellCompDirective) {
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
		},
	}

	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all decrypted file")
	cmd.Flags().BoolVarP(&flagClip, "clip", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().BoolVarP(&flagQR, "qr", "q", false, "Generate a QR code for the decrypted password")
	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Filter by field name...")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	cmd.MarkFlagsMutuallyExclusive("all", "field")
	cmd.MarkFlagsMutuallyExclusive("all", "qr")
	cmd.MarkFlagsMutuallyExclusive("qr", "field")
	cmd.MarkFlagsMutuallyExclusive("qr", "clip")

	cmd.RegisterFlagCompletionFunc("field", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		fields := []string{
			"email",
			"host",
			"notes",
			"otp",
			"pass",
			"password",
			"port",
			"recovery_codes",
			"secret_key",
			"url",
			"user",
			"username",
			"aws.region",
			"aws.account_id",
			"aws.access_key",
			"aws.secret_access_key",
		}

		var out []string
		for _, f := range fields {
			if strings.HasPrefix(f, toComplete) {
				out = append(out, f)
			}
		}
		return out, cobra.ShellCompDirectiveNoFileComp
	})

	return cmd
}

func Run(cmd *cobra.Command, path string) {
	var v string
	var b = openpgp.Decrypt(term.ReadPassword(flagPassphrase), path)
	var c = card.New(b)

	switch flagField {
	case "email":
		v = c.Email
	case "host":
		v = c.Host
	case "notes":
		v = c.Notes
	case "otp":
		v = otp.Get(c.OTP)
	case "pass":
		v = c.Password
	case "password":
		v = c.Password
	case "port":
		v = c.Port
	case "recovery_codes":
		v = c.RecoveryCodes
	case "secret_key":
		v = c.SecretKey
	case "url":
		v = c.URL
	case "user":
		v = c.Username
	case "username":
		v = c.Username
	case "aws.region":
		v = c.AWS.Region
	case "aws.account_id":
		v = c.AWS.AccountId
	case "aws.access_key":
		v = c.AWS.AccessKey
	case "aws.secret_access_key":
		v = c.AWS.SecretAccessKey
	default:
		if flagAll {
			v = b
		} else {
			v = c.Password
		}
	}

	if flagClip {
		clipboard.Write(v)
	} else if flagQR {
		qr.Generate(v)
	} else {
		fmt.Fprintln(cmd.OutOrStdout(), v)
	}
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
