package show

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

var clip bool
var field string
var passphrase string

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show [flags] path/to/file",
		Short: "Show and decrypt selected data. By default, only the password field is shown.",
    	Example: "  pm show -p <passphrase> -c wifi/theforce\n" +
    			 "  pm show -p <passphrase> -c -f aws.access_key com/aws",
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
	cmd.Flags().StringVarP(&field, "field", "f", "", "Filter by field name. The default is <password>.\nAllowed fields:\n host | otp | pass | password | port | url | user | username\n aws.region | aws.account_id | aws.access_key | aws.secret_access_key")

	return cmd
}

func Run(cmd *cobra.Command, path string) {
	var v string
	var c = card.New(openpgp.Decrypt(passphrase, path))

	switch field {
	case "host":
		v = c.Host
	case "otp":
		v = c.OTP
	case "pass":
		v = c.Password
	case "password":
		v = c.Password
	case "port":
		v = c.Port
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
		v = c.Password
	}

	if clip {
		clipboard.Write(v)
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
