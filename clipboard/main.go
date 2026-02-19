package clipboard

import (
	"errors"
	"os"
	"runtime"

	"golang.design/x/clipboard"
)

// Write copies text into the system clipboard.
func Write(text string) error {
	if runtime.GOOS == "linux" && os.Getenv("DISPLAY") == "" {
		return errors.New("Clipboard disabled: no X11 display detected")
	}

	if err := clipboard.Init(); err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtText, []byte(text))

	return nil
}
