package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const version = "0.1.0"

type cuiApp struct {
	Footer *tview.TextView

	MethodDropdown *tview.DropDown
	UrlInput       *tview.InputField

	Request             *tview.Flex
	RequestKindDropdown *tview.DropDown
	RequestBody         *tview.TextArea
	RequestFormData     *tview.Table
	RequestHeaders      *tview.Table
	RequestParameters   *tview.TextView

	Response        *tview.Flex
	ResponseStatus  *tview.TextView
	ResponseBody    *tview.TextView
	ResponseHeaders *tview.Table
}

func main() {
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

	// no help, because this is a hidden debug feature
	serve := flag.Bool("server", false, "start simple echo dev server")
	flag.Parse()

	if *serve {
		cuiDevServer(methods)
		return
	}

	logfile, err := os.OpenFile("cui.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.SetFlags(1 | 2)

	app := tview.NewApplication()
	hasResponse := false
	responseView := "body"
	requestView := "RequestBody"

	requestKinds := []string{"Form Data", "JSON", "Raw"}
	requestKind := 2 // requestKinds is zero-indexed

	cui := cuiApp{
		Footer:              tview.NewTextView(),
		MethodDropdown:      tview.NewDropDown(),
		UrlInput:            tview.NewInputField(),
		Request:             tview.NewFlex(),
		RequestKindDropdown: tview.NewDropDown(),
		RequestBody:         tview.NewTextArea(),
		RequestFormData:     tview.NewTable(),
		RequestHeaders:      tview.NewTable(),
		RequestParameters:   tview.NewTextView(),
		Response:            tview.NewFlex(),
		ResponseStatus:      tview.NewTextView(),
		ResponseBody:        tview.NewTextView(),
		ResponseHeaders:     tview.NewTable(),
	}

	req := cuiRequest{
		Method:     http.MethodGet,
		URL:        "http://example.com",
		Headers:    make(map[string]string),
		Parameters: make(map[string]string),
		Body:       "",
	}

	setInstructions(&cui, "")
	cui.MethodDropdown.SetOptions(methods, nil).SetCurrentOption(methodGet)
	cui.UrlInput.SetLabel("URL: ").SetPlaceholder("http://example.com")

	cui.RequestKindDropdown.SetOptions(requestKinds, nil).SetCurrentOption(requestKind)

	methodAndUrl := tview.NewFlex().
		AddItem(cui.MethodDropdown, 10, 0, false).
		AddItem(cui.UrlInput, 0, 1, false)

	cui.Response.SetDirection(tview.FlexRow).
		AddItem(cui.ResponseStatus, 1, 0, false).
		AddItem(cui.ResponseBody, 0, 1, true)

	cui.Request.SetDirection(tview.FlexRow).
		AddItem(cui.RequestKindDropdown, 1, 0, false).
		AddItem(cui.RequestBody, 0, 1, true)

	newRequest := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(methodAndUrl, 1, 0, false).
		AddItem(cui.Request, 0, 1, false).
		AddItem(cui.Response, 0, 1, false)

	newRequest.SetBorder(true).SetTitle(" New Request ")

	inner := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" Request History "), 0, 1, false).
		AddItem(newRequest, 0, 3, false)

	header := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(fmt.Sprintf("cUI v%s", version))

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

	cui.RequestKindDropdown.SetDoneFunc(func(key tcell.Key) {
		// TODO: this leaves the dropdown focused...
		app.SetFocus(cui.Request)
		setInstructions(&cui, "RequestBody")

		_, kind := cui.RequestKindDropdown.GetCurrentOption()
		if kind == "Raw" {
			delete(req.Headers, "Content-Type")
			// TODO: ensure we have the raw text entry for body
		} else if kind == "JSON" {
			req.Headers["Content-Type"] = "application/json"
			// TODO: ensure we have the raw text entry for body
		} else {
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
			// TODO: we need to set the key/val form entry for body
		}
	})
	cui.RequestKindDropdown.SetSelectedFunc(func(text string, index int) {
		app.SetFocus(cui.Request)
		setInstructions(&cui, "RequestBody")

		_, kind := cui.RequestKindDropdown.GetCurrentOption()
		if kind == "Raw" {
			delete(req.Headers, "Content-Type")
			// TODO: ensure we have the raw text entry for body
		} else if kind == "JSON" {
			req.Headers["Content-Type"] = "application/json"
			// TODO: ensure we have the raw text entry for body
		} else {
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
			// TODO: we need to set the key/val form entry for body
		}
	})

	cui.RequestBody.SetChangedFunc(func() {
		req.Body = cui.RequestBody.GetText()
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
				return nil // return nil to prevent u from being inserted
			} else if event.Rune() == 114 && hasResponse { // r
				if responseView == "body" {
					setInstructions(&cui, "ResponseBody")
				} else {
					setInstructions(&cui, "ResponseHeaders")
				}
				app.SetFocus(cui.Response)
			} else if event.Rune() == 101 { // e
				setInstructions(&cui, requestView)
				app.SetFocus(cui.Request)
				return nil // return nil to prevent e from being inserted
			} else if event.Key() == tcell.KeyEnter {
				if err := sendRequest(req, &cui, &hasResponse); err != nil {
					panic(err)
				}

				setInstructions(&cui, "ResponseBody")
				responseView = "body"
				app.SetFocus(cui.Response)
			}
		}

		if app.GetFocus() == cui.ResponseBody {
			if event.Key() == tcell.KeyEscape {
				setInstructions(&cui, "WithResponse")
				app.SetFocus(main)
			} else if event.Rune() == 116 { // t
				setInstructions(&cui, "ResponseHeaders")
				responseView = "headers"

				cui.Response.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.ResponseHeaders, 0, 1, true)

				app.SetFocus(cui.Response)

				return nil
			}

		}

		if app.GetFocus() == cui.ResponseHeaders {
			if event.Key() == tcell.KeyEscape {
				setInstructions(&cui, "WithResponse")
				app.SetFocus(main)
			} else if event.Rune() == 116 { // t
				setInstructions(&cui, "ResponseBody")
				responseView = "body"

				cui.Response.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.ResponseStatus, 1, 0, false).
					AddItem(cui.ResponseBody, 0, 1, true)

				app.SetFocus(cui.Response)

				return nil
			}
		}

		if app.GetFocus() == cui.RequestBody {
			if event.Key() == tcell.KeyEscape {
				if hasResponse {
					setInstructions(&cui, "WithResponse")
				} else {
					setInstructions(&cui, "")
				}
				app.SetFocus(main)
			} else if event.Key() == tcell.KeyCtrlH {
				requestView = "RequestHeaders"
				setInstructions(&cui, requestView)

				cui.RequestHeaders.Clear()
				cui.Request.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.RequestHeaders, 0, 1, true)

				i := 0
				for header, value := range req.Headers {
					cui.RequestHeaders.SetCell(i, 0, tview.NewTableCell(header))
					cui.RequestHeaders.SetCell(i, 1, tview.NewTableCell(value))
					i += 1
				}

				app.SetFocus(cui.Response)

				return nil
			} else if event.Key() == tcell.KeyCtrlP {
				requestView = "RequestParameters"
				setInstructions(&cui, requestView)
			} else if event.Key() == tcell.KeyCtrlK {
				setInstructions(&cui, "RequestKindDropdown")
				app.SetFocus(cui.RequestKindDropdown)
			}
		}

		if app.GetFocus() == cui.RequestHeaders {
			if event.Key() == tcell.KeyEscape {
				if hasResponse {
					setInstructions(&cui, "WithResponse")
				} else {
					setInstructions(&cui, "")
				}
				app.SetFocus(main)
			} else if event.Rune() == 97 { // a
			} else if event.Rune() == 100 { // d
			} else if event.Rune() == 98 { // b
			} else if event.Rune() == 112 { // p
			}
		}

		if app.GetFocus() == cui.RequestParameters {
			if event.Key() == tcell.KeyEscape {
				if hasResponse {
					setInstructions(&cui, "WithResponse")
				} else {
					setInstructions(&cui, "")
				}
				app.SetFocus(main)
			} else if event.Rune() == 97 { // a
			} else if event.Rune() == 100 { // d
			} else if event.Rune() == 98 { // b
			} else if event.Rune() == 104 { // h
			}
		}

		return event
	})

	if err := app.SetRoot(main, true).Run(); err != nil {
		panic(err)
	}
}
