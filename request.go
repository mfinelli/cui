package main

import (
	"io/ioutil"
	"net/http"

	"github.com/rivo/tview"
)

type cuiRequest struct {
	Method string
	URL string
}

func sendRequest(req cuiRequest, prim *tview.TextView) error {
	client := &http.Client{}

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

	prim.SetText(string(body))

	return nil
}
