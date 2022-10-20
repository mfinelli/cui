package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rivo/tview"
)

type cuiRequest struct {
	Method string
	URL    string
}

func sendRequest(req cuiRequest, cui *cuiApp, hasResponse *bool) error {
	client := &http.Client{}
	cui.ResponseBody.Clear()
	cui.ResponseHeaders.Clear()

	r, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	cui.ResponseStatus.SetText(fmt.Sprintf("Status: %d", res.StatusCode))
	cui.ResponseBody.SetText(string(body)).ScrollToBeginning()

	cui.ResponseHeaders.SetCell(0, 0, tview.NewTableCell("Header"))
	cui.ResponseHeaders.SetCell(0, 1, tview.NewTableCell("Value"))

	i := 1
	for k, v := range res.Header {
		for _, vv := range v {
			cui.ResponseHeaders.SetCell(i, 0, tview.NewTableCell(k))
			cui.ResponseHeaders.SetCell(i, 1, tview.NewTableCell(vv))
			i += 1
		}
	}

	*hasResponse = true

	return nil
}
