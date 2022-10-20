package main

import (
	// "github.com/rivo/tview"
)

func setInstructions(cui *cuiApp, instr string) {
	if instr == "MethodDropdown" {
		cui.Footer.SetText(" (esc) cancel  (enter) confirm  (↑/↓) navigate")
	} else if instr == "UrlInput" {
		cui.Footer.SetText(" (enter) finish entering text")
	} else {
		cui.Footer.SetText(" (q) quit  (m) set method  (u) set url")
	}
}
