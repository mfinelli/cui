package main

import (
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type cuiApp struct {
	Footer *tview.TextView

	MethodDropdown *tview.DropDown
	UrlInput *tview.InputField

	Response *tview.Flex
	ResponseStatus *tview.TextView
	ResponseBody *tview.TextView
	ResponseHeaders *tview.Table
}

func main() {
	app := tview.NewApplication()

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

	cui := cuiApp {
		Footer: tview.NewTextView(),
		MethodDropdown: tview.NewDropDown(),
		UrlInput: tview.NewInputField(),
		Response: tview.NewFlex(),
		ResponseStatus: tview.NewTextView(),
		ResponseBody: tview.NewTextView(),
		ResponseHeaders: tview.NewTable(),
	}

	req := cuiRequest{
		Method: http.MethodGet,
		URL: "http://example.com",
	}

	setInstructions(&cui, "")
	cui.MethodDropdown.SetOptions(methods, nil).SetCurrentOption(methodGet)
	cui.UrlInput.SetLabel("URL: ").SetPlaceholder("http://example.com")

	methodAndUrl := tview.NewFlex().
		AddItem(cui.MethodDropdown, 10, 0, false).
		AddItem(cui.UrlInput, 0, 1, false)

	cui.Response.SetDirection(tview.FlexRow).
		AddItem(cui.ResponseStatus, 1, 0, false).
		AddItem(cui.ResponseBody, 0, 1, false)

	newRequest := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(methodAndUrl, 1, 0, false).
		AddItem(tview.NewTextView().SetText("request"), 0, 1, false).
		AddItem(cui.Response, 0, 1, false)

	newRequest.SetBorder(true).SetTitle(" New Request ")

	inner := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" Request History "), 0, 1, false).
		AddItem(newRequest, 0, 3, false)

	header := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("cUI v1.0.0")

	main := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(inner, 0, 1, false).
		AddItem(cui.Footer, 1, 0, false)

	cui.MethodDropdown.SetDoneFunc(func(key tcell.Key) {
		// TODO: this leaves the dropdown focused...
		app.SetFocus(main)
		setInstructions(&cui, "")
		_, req.Method = cui.MethodDropdown.GetCurrentOption()
	})
	cui.MethodDropdown.SetSelectedFunc(func(text string, index int) {
		app.SetFocus(main)
		setInstructions(&cui, "")
		_, req.Method = cui.MethodDropdown.GetCurrentOption()
	})
	cui.UrlInput.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(main)
		setInstructions(&cui, "")
		req.URL = cui.UrlInput.GetText()
	})

	// fmt.Printf("%s: %d", string('x'), int('x'))
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if app.GetFocus() == main { // we're not focused on anything
			if event.Rune() == 113 { // q
				app.Stop()
			} else if event.Rune() == 109 { // m
				setInstructions(&cui, "MethodDropdown")
				app.SetFocus(cui.MethodDropdown)
			} else if event.Rune() == 117 { // u
				setInstructions(&cui, "UrlInput")
				app.SetFocus(cui.UrlInput)
				return nil
			} else if event.Key() == tcell.KeyEnter {
				if err := sendRequest(req, &cui); err != nil {
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
