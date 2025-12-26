package file

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/explorer"
	"github.com/nicola-strappazzon/pm/file"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/path"
	"github.com/nicola-strappazzon/pm/term"

	"github.com/spf13/cobra"
)

var fileName string
var flagDelete string
var flagExtract string
var flagInclude string
var flagList bool
var flagOutput string
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "file",
		Short: "Manage files stored inside a GPG-encrypted container.",
		Example: "  pm file <TAB>\n" +
			"  pm file passport -p <passphrase> -l\n" +
			"  pm file passport -p <passphrase> -i /path/file/to/front.png\n" +
			"  pm file passport -p <passphrase> -e front.png -o front.png\n" +
			"  pm file passport -p <passphrase> -d front.png\n",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagList, "list", "l", false, "List files stored in the container.")
	cmd.Flags().StringVarP(&flagDelete, "delete", "d", "", "Delete a file from container.")
	cmd.Flags().StringVarP(&flagExtract, "extract", "e", "", "Extract a file from container.")
	cmd.Flags().StringVarP(&flagInclude, "include", "i", "", "Include a file into container.")
	cmd.Flags().StringVarP(&flagOutput, "output", "o", "", "Output path to extract file.")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file.")

	cmd.MarkFlagsMutuallyExclusive("extract", "include", "list", "delete")
	cmd.MarkFlagsMutuallyExclusive("include", "list", "delete", "output")

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var tmpCard card.Card
	var pathCard string = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.IsNotFile() {
		explorer.PrintTree(p.Absolute())
		return nil
	}

	if flagInclude != "" {
		fileName = file.Name(flagInclude)
	}

	if flagExtract != "" {
		fileName = file.Name(flagExtract)
	}

	if flagDelete != "" {
		fileName = flagDelete
	}

	tmpCard = card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		p.Full(),
	))
	tmpCard.Path = p.Full()

	if flagInclude != "" {
		if tmpCard.Files.Exist(card.File{Name: fileName}) {
			return fmt.Errorf("File %s already exists; operation aborted.", fileName)
		}

		tmpCard.Files.Add((&card.File{}).Load(flagInclude))
		tmpCard.Save()

		fmt.Fprintf(cmd.OutOrStdout(), "Added file %s to the GPG-encrypted container %s.\n", fileName, p.Path())
		return nil
	}

	if flagExtract != "" {
		if tmpCard.Files.Exist(card.File{Name: fileName}) {
			tmpCard.Files.Get(card.File{Name: fileName}).Save(flagOutput)
			fmt.Fprintf(cmd.OutOrStdout(), "Saved file %s to %s.\n", fileName, flagOutput)
		} else {
			return fmt.Errorf("File %s not exists; operation aborted.", fileName)
		}
	}

	if flagDelete != "" {
		if tmpCard.Files.Exist(card.File{Name: fileName}) {
			tmpCard.Files.Delete(card.File{Name: fileName})
			tmpCard.Save()
			fmt.Fprintf(cmd.OutOrStdout(), "Deleted file %s from the GPG-encrypted container %s.\n", fileName, p.Path())
		} else {
			return fmt.Errorf("File %s not exists; operation aborted.", fileName)
		}
	}

	return nil
}
