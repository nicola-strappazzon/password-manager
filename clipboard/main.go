package clipboard

import (
	"log"
	"os"
	"runtime"

	"golang.design/x/clipboard"
)

var available bool

func Init() {
	if runtime.GOOS == "linux" && os.Getenv("DISPLAY") == "" {
		log.Println("Clipboard disabled: no X11 display detected")
		return
	}

	if err := clipboard.Init(); err != nil {
		log.Fatal(err.Error())
	}

	available = true
}

// Write copies text into the system clipboard.
func Write(text string) {
	if !available {
		log.Println("Clipboard not available")
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(text))
}
