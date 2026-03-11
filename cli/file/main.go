package file

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"
	"github.com/nicola-strappazzon/password-manager/internal/file"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/dustin/go-humanize"
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
		Short: "Manage files stored in encrypted containers",
		Example: "  pm file <TAB>\n" +
			"  pm file passport -p <passphrase> -l\n" +
			"  pm file passport -p <passphrase> -i /path/file/to/front.png\n" +
			"  pm file passport -p <passphrase> -e front.png -o front.png\n" +
			"  pm file passport -p <passphrase> -d front.png\n",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagList, "list", "l", false, "List files stored in the container")
	cmd.Flags().StringVarP(&flagDelete, "delete", "d", "", "Delete a file from container")
	cmd.Flags().StringVarP(&flagExtract, "extract", "e", "", "Extract a file from container")
	cmd.Flags().StringVarP(&flagInclude, "include", "i", "", "Include a file into container")
	cmd.Flags().StringVarP(&flagOutput, "output", "o", "", "Output path to extract file")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	cmd.MarkFlagsMutuallyExclusive("extract", "include", "list", "delete")
	cmd.MarkFlagsMutuallyExclusive("include", "list", "delete", "output")

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var tmpCard card.Card
	var pathCard string = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.IsDirectory() {
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

		cmd.Printf("Added file %s to the GPG-encrypted container %s.\n", fileName, p.Path())
		return nil
	}

	if flagExtract != "" {
		if tmpCard.Files.Exist(card.File{Name: fileName}) {
			tmpCard.Files.Get(card.File{Name: fileName}).Save(flagOutput)
			cmd.Printf("Saved file %s to %s.\n", fileName, flagOutput)
		} else {
			return fmt.Errorf("File %s does not exist; operation aborted.", fileName)
		}
	}

	if flagDelete != "" {
		if tmpCard.Files.Exist(card.File{Name: fileName}) {
			tmpCard.Files.Delete(card.File{Name: fileName})
			tmpCard.Save()
			cmd.Printf("Deleted file %s from the GPG-encrypted container %s.\n", fileName, p.Path())
		} else {
			return fmt.Errorf("File %s does not exist; operation aborted.", fileName)
		}
	}

	if flagList {
		cmd.Printf("Files inside %s:\n", p.Path())

		for _, file := range tmpCard.Files {
			cmd.Printf(" - %s (%s)\n", file.Name, humanize.Bytes(file.Size()))
		}
	}

	return nil
}
