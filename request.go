package main

import (
	"io/ioutil"
	"net/http"

	// "github.com/rivo/tview"
)

type cuiRequest struct {
	Method string
	URL string
}

func sendRequest(req cuiRequest, cui *cuiApp) error {
	client := &http.Client{}
	cui.ResponseBody.Clear()

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

	cui.ResponseBody.SetText(string(body)).ScrollToBeginning()

	return nil
}
