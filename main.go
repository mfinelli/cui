package main

import (
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	req := cuiRequest{
		Method: http.MethodGet,
		URL: "http://example.com",
	}

	methods := []string{
		http.MethodDelete,
		http.MethodHead,
		http.MethodGet,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
	}
	methodGet := 2 // methods is zero-indexed

	methodDropdown := tview.NewDropDown().SetOptions(methods, nil).SetCurrentOption(methodGet)
	urlInput := tview.NewInputField().SetLabel("URL: ").SetPlaceholder("http://example.com")

	methodAndUrl := tview.NewFlex().
		AddItem(methodDropdown, 10, 0, false).
		AddItem(urlInput, 0, 1, false)

	respBody := tview.NewTextView()

	newRequest := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(methodAndUrl, 1, 0, false).
		AddItem(tview.NewTextView().SetText("request"), 0, 1, false).
		AddItem(respBody.SetText("response"), 0, 1, false)

	newRequest.SetBorder(true).SetTitle(" New Request ")

	inner := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" Request History "), 0, 1, false).
		AddItem(newRequest, 0, 3, false)

	header := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("cUI v1.0.0")

	main := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(inner, 0, 1, false).
		AddItem(tview.NewTextView().SetText(" (q) quit  (m) set method  (u) set url"), 1, 0, false)

	methodDropdown.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(main)
		_, req.Method = methodDropdown.GetCurrentOption()
	})
	methodDropdown.SetSelectedFunc(func(text string, index int) {
		app.SetFocus(main)
		_, req.Method = methodDropdown.GetCurrentOption()
	})
	urlInput.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(main)
		req.URL = urlInput.GetText()
	})

	// fmt.Printf("%s: %d", string('x'), int('x'))
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if app.GetFocus() == main { // we're not focused on anything
			if event.Rune() == 113 { // q
				app.Stop()
			} else if event.Rune() == 109 { // m
				app.SetFocus(methodDropdown)
			} else if event.Rune() == 117 { // u
				app.SetFocus(urlInput)
				return nil
			} else if event.Key() == tcell.KeyEnter {
				if err := sendRequest(req, respBody); err != nil {
					panic(err)
				}
			}
		}

		return event
	})

	if err := app.SetRoot(main, true).Run(); err != nil {
		panic(err)
	}
}
