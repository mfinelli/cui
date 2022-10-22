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
		cui.Footer.SetText(" (q) quit  (enter) send request  (m) set method  (u) set url  (e) edit request  (h) edit request headers  (p) edit request parameters  (r) focus response  (c) clear")
	} else if instr == "RequestBody" {
		cui.Footer.SetText(" (esc) done  (crtl^k) set kind  (ctrl^e) edit in $EDITOR")
	} else if instr == "RequestHeaders" {
		cui.Footer.SetText(" (esc) done  (a) add new header  (d) delete selected header")
	} else if instr == "RequestParameters" {
		cui.Footer.SetText(" (esc) done  (a) add query parameter  (d) deleted selected parameter")
	} else {
		cui.Footer.SetText(" (q) quit  (enter) send request  (m) set method  (u) set url  (e) edit request body  (h) edit request headers  (p) edit request parameters  (c) clear")
	}
}
