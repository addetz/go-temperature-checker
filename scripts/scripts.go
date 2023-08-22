package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/addetz/railway-go-demo/apis"
	"honnef.co/go/js/dom"
)

//go:embed cities
var citiesInput string

func main() {
	cities := strings.Split(citiesInput, "\n")
	document := dom.GetWindow().Document()
	populateCities(document, cities)

	selection := document.GetElementByID("citiesDropdown")
	selection.AddEventListener("change", false, func(e dom.Event) {
		selEl := selection.(*dom.HTMLSelectElement)
		index := selEl.SelectedIndex
		selectedText := selEl.Options()[index].Text
		city := strings.Split(selectedText, ",")[0]
		go func(callback func(bs *apis.BackendResponse)) {
			resp, err := http.Get(fmt.Sprintf("/weather/%s", city))
			if err != nil {
				log.Fatal(err)
			}
			bs, err := apis.NewBackendResponse(resp)
			if err != nil {
				log.Fatal(err)
			}
			callback(bs)

		}(weatherCallback)
	})
}

func populateCities(document dom.Document, cities []string) {
	selEl := document.GetElementByID("citiesDropdown").(*dom.HTMLSelectElement)
	for _, c := range cities {
		o := document.CreateElement("option")
		o.SetTextContent(c)
		selEl.AppendChild(o)
	}
}

func weatherCallback(bs *apis.BackendResponse) {
	log.Println(*bs)
}
