package main

func setInstructions(cui *cuiApp, instr string) {
	if instr == "MethodDropdown" {
		cui.Footer.SetText(" (esc) cancel  (enter) confirm  (↑/↓) navigate")
	} else if instr == "UrlInput" {
		cui.Footer.SetText(" (enter) finish entering text")
	} else if instr == "ResponseBody" {
		cui.Footer.SetText(" (q) lose focus  (h/j/k/l) navigate  (s) save response body to file  (t) response headers")
	} else if instr == "ResponseHeaders" {
		cui.Footer.SetText(" (q) lose focus  (t) response body")
	} else if instr == "WithResponse" {
		cui.Footer.SetText(" (q) quit  (m) set method  (u) set url  (r) focus response")
	} else {
		cui.Footer.SetText(" (q) quit  (m) set method  (u) set url")
	}
}
