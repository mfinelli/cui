// cui: http request/response tui
// Copyright 2022 Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
)

func setInstructions(cui *cuiApp, instr string) {
	if instr == "RequestHistory" {
		cui.FooterInstruction.SetText(" (esc) done  (enter) load  (↑/↓) navigate  (d) delete)")
	} else if instr == "MethodDropdown" {
		cui.FooterInstruction.SetText(" (esc) cancel  (enter) confirm  (↑/↓) navigate")
	} else if instr == "RequestKindDropdown" {
		cui.FooterInstruction.SetText(" (esc) cancel  (enter) confirm  (↑/↓) navigate")
	} else if instr == "UrlInput" {
		cui.FooterInstruction.SetText(" (enter) finish entering text")
	} else if instr == "ResponseBody" {
		cui.FooterInstruction.SetText(" (esc) done  (h/j/k/l) navigate  (s) save response body to file  (t) response headers")
	} else if instr == "ResponseHeaders" {
		cui.FooterInstruction.SetText(" (esc) done  (t) response body")
	} else if instr == "RequestBodyTextarea" {
		cui.FooterInstruction.SetText(" (esc) done  (crtl^k) set kind  (ctrl^h) edit headers (ctrl^p) edit query parameters (ctrl^e) edit in $EDITOR")
	} else if instr == "RequestBodyFormdata" {
		cui.FooterInstruction.SetText(" TODO")
	} else if instr == "RequestHeaders" {
		cui.FooterInstruction.SetText(" (esc) done  (↑/↓) cycle headers  (a) add new header  (d) delete selected header  (b) edit request body  (p) edit query parameters")
	} else if instr == "RequestHeaderAdd" {
		cui.FooterInstruction.SetText(" (esc) cancel  (enter) done  (tab) cycle key/value")
	} else if instr == "RequestParameters" {
		cui.FooterInstruction.SetText(" (esc) done  (a) add query parameter  (d) deleted selected parameter  (b) edit request body  (h) edit headers")
	} else if instr == "SaveResponse" {
		cui.FooterInstruction.SetText(" (esc) cancel  (enter) done ")

	} else if instr == "FileExists" {
		cui.FooterInstruction.SetText("(!) File With this name already  exists (w) replace (esc) cancel")

	} else {
		if instr != "" {
			log.Printf("Couldn't find instructions for '%s'\n", instr)
		}

		if cui.ViewHasResponse {
			cui.FooterInstruction.SetText(" (q) quit  (enter) send request  (m) set method  (u) set url  (e) edit request  (r) focus response  (h) history  (c) clear")
		} else {
			cui.FooterInstruction.SetText(" (q) quit  (enter) send request  (m) set method  (u) set url  (e) edit request  (h) history  (c) clear")
		}
	}
}
