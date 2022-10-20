package main

import (
	// "github.com/rivo/tview"
)

func setInstructions(cui *cuiApp, instr string) {
	if instr == "MethodDropdown" {
		cui.Footer.SetText(" (esc) cancel  (enter) confirm  (↑/↓) navigate")
	} else if instr == "UrlInput" {
		cui.Footer.SetText(" (enter) finish entering text")
	} else if instr == "ResponseBody" {
		cui.Footer.SetText(" (q) lose focus  (h/j/k/l) navigate  (s) save response body to file  (h) response headers")
	} else if instr == "ResponseHeaders" {
		cui.Footer.SetText(" (q) lose focus  (b) response body")
	} else {
		cui.Footer.SetText(" (q) quit  (m) set method  (u) set url")
	}
}
