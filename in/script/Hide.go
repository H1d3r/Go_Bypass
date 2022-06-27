package script

import (
	"github.com/gonutz/ide/w32"
)

//隐藏console
func Hide() {
	ShowConsoleAsync(w32.SW_HIDE)
}

func ShowConsoleAsync(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}
