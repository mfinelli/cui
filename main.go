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
	RequestHeaderKey    *tview.InputField
	RequestHeaderValue  *tview.InputField
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
	requestInputType := "Textarea"

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
		RequestHeaderKey:    tview.NewInputField(),
		RequestHeaderValue:  tview.NewInputField(),
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

	setInstructions(&cui, "", hasResponse)
	cui.MethodDropdown.SetOptions(methods, nil).SetCurrentOption(methodGet)
	cui.UrlInput.SetLabel("URL: ").SetPlaceholder("http://example.com")
	cui.RequestHeaderKey.SetLabel("Key: ")
	cui.RequestHeaderValue.SetLabel("Value: ")

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
		setInstructions(&cui, "", hasResponse)
		_, req.Method = cui.MethodDropdown.GetCurrentOption()
	})
	cui.MethodDropdown.SetSelectedFunc(func(text string, index int) {
		app.SetFocus(main)
		setInstructions(&cui, "", hasResponse)
		_, req.Method = cui.MethodDropdown.GetCurrentOption()
	})
	cui.UrlInput.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(main)
		setInstructions(&cui, "", hasResponse)
		req.URL = cui.UrlInput.GetText()
	})

	cui.RequestKindDropdown.SetDoneFunc(func(key tcell.Key) {
		// TODO: this leaves the dropdown focused...
		app.SetFocus(cui.Request)

		_, kind := cui.RequestKindDropdown.GetCurrentOption()
		if kind == "Raw" {
			requestInputType = "Textarea"
			delete(req.Headers, "Content-Type")
			// TODO: ensure we have the raw text entry for body
		} else if kind == "JSON" {
			requestInputType = "Textarea"
			req.Headers["Content-Type"] = "application/json"
			// TODO: ensure we have the raw text entry for body
		} else {
			requestInputType = "Formdata"
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
			// TODO: we need to set the key/val form entry for body
		}

		setInstructions(&cui, requestView+requestInputType, hasResponse)
	})
	cui.RequestKindDropdown.SetSelectedFunc(func(text string, index int) {
		app.SetFocus(cui.Request)

		_, kind := cui.RequestKindDropdown.GetCurrentOption()
		if kind == "Raw" {
			requestInputType = "Textarea"
			delete(req.Headers, "Content-Type")
			// TODO: ensure we have the raw text entry for body
		} else if kind == "JSON" {
			requestInputType = "Textarea"
			req.Headers["Content-Type"] = "application/json"
			// TODO: ensure we have the raw text entry for body
		} else {
			requestInputType = "Formdata"
			req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
			// TODO: we need to set the key/val form entry for body
		}

		setInstructions(&cui, requestView+requestInputType, hasResponse)
	})

	cui.RequestBody.SetChangedFunc(func() {
		req.Body = cui.RequestBody.GetText()
	})

	// fmt.Printf("%s: %d", string('x'), int('x'))
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		focus := app.GetFocus()

		if event.Key() == tcell.KeyEnter {
			if focus == main {
				if err := sendRequest(req, &cui, &hasResponse); err != nil {
					panic(err)
				}

				responseView = "body"
				setInstructions(&cui, "ResponseBody", hasResponse)
				app.SetFocus(cui.Response)
			}
		} else if event.Key() == tcell.KeyEscape {
			if focus == cui.ResponseBody || focus == cui.ResponseHeaders || focus == cui.RequestBody || focus == cui.RequestHeaders || focus == cui.RequestParameters {
				setInstructions(&cui, "", hasResponse)
				app.SetFocus(main)
			} else if focus == cui.RequestHeaderKey || focus == cui.RequestHeaderValue {
				setInstructions(&cui, "RequestHeaders", hasResponse)
				setEditHeadersPlain(&cui, &req)
				app.SetFocus(cui.Request)
			}
		} else if event.Key() == tcell.KeyTab {
			if focus == cui.RequestHeaderKey {
				app.SetFocus(cui.RequestHeaderValue)
			} else if focus == cui.RequestHeaderValue {
				app.SetFocus(cui.RequestHeaderKey)
			}
		} else if event.Key() == tcell.KeyCtrlH {
			if focus == cui.RequestBody {
				requestView = "RequestHeaders"
				setInstructions(&cui, requestView, hasResponse)
				setEditHeadersPlain(&cui, &req)
				app.SetFocus(cui.Request)
			}
		} else if event.Key() == tcell.KeyCtrlK {
			if focus == cui.RequestBody {
				setInstructions(&cui, "RequestKindDropdown", hasResponse)
				app.SetFocus(cui.RequestKindDropdown)
			}
		} else if event.Key() == tcell.KeyCtrlP {
			if focus == cui.RequestBody {
				setInstructions(&cui, "RequestParameters", hasResponse)
				// TODO: set param view
				app.SetFocus(cui.Request)
			}
		} else if event.Rune() == 97 { // a
			if focus == cui.RequestHeaders {
				setInstructions(&cui, "RequestHeaderAdd", hasResponse)
				setEditHeadersAdd(&cui, &req)
				app.SetFocus(cui.Request)
				return nil // prevent "a" from being entered:w
			}
		} else if event.Rune() == 98 { // b
			if focus == cui.RequestHeaders || focus == cui.RequestParameters {
				requestView = "RequestBody"
				setInstructions(&cui, requestView+requestInputType, hasResponse)

				cui.Request.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.RequestKindDropdown, 1, 0, false).
					AddItem(cui.RequestBody, 0, 1, true)
				app.SetFocus(cui.Request)
				return nil // prevent "b" from being entered
			}
		} else if event.Rune() == 100 { // d
		} else if event.Rune() == 101 { // e
			if focus == main {
				if requestView == "RequestBody" {
					setInstructions(&cui, requestView+requestInputType, hasResponse)
				} else {
					setInstructions(&cui, requestView, hasResponse)
				}
				app.SetFocus(cui.Request)
				return nil // prevent "e" from being inserted
			}
		} else if event.Rune() == 104 { // h
		} else if event.Rune() == 109 { // m
			if focus == main {
				setInstructions(&cui, "MethodDropdown", hasResponse)
				app.SetFocus(cui.MethodDropdown)
			}
		} else if event.Rune() == 112 { // p
		} else if event.Rune() == 113 { // q
			if focus == main {
				app.Stop()
				return nil
			}
		} else if event.Rune() == 114 { // r
			if focus == main && hasResponse {
				if responseView == "body" {
					setInstructions(&cui, "ResponseBody", hasResponse)
				} else {
					setInstructions(&cui, "ResponseHeaders", hasResponse)
				}

				app.SetFocus(cui.Response)
			}
		} else if event.Rune() == 116 { // t
			if focus == cui.ResponseBody {
				responseView = "headers"
				setInstructions(&cui, "ResponseHeaders", hasResponse)

				cui.Response.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.ResponseHeaders, 0, 1, true)
				app.SetFocus(cui.Response)
			} else if focus == cui.ResponseHeaders {
				responseView = "body"
				setInstructions(&cui, "ResponseBody", hasResponse)

				cui.Response.Clear().SetDirection(tview.FlexRow).
					AddItem(cui.ResponseStatus, 1, 0, false).
					AddItem(cui.ResponseBody, 0, 1, true)
				app.SetFocus(cui.Response)
			}
		} else if event.Rune() == 117 { // u
			if focus == main {
				setInstructions(&cui, "UrlInput", hasResponse)
				app.SetFocus(cui.UrlInput)
				return nil // prevent "u" from being entered
			}
		}

		return event
	})

	if err := app.SetRoot(main, true).Run(); err != nil {
		panic(err)
	}
}
