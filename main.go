package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"text/template"
)

func translate() {
	key := ""
	endpoint := "https://api.cognitive.microsofttranslator.com/"
	uri := endpoint + "/translate?api-version=3.0"
	location := "eastus"

	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("from", "en")
	q.Add("to", "es")
	q.Add("to", "it")
	q.Add("to", "de")
	q.Add("to", "id")
	q.Add("to", "pt")
	u.RawQuery = q.Encode()

	body := []struct {
		Text string
	}{
		{Text: "Hello friend! What did you do today?"},
	}
	b, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", key)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("%s\n", prettyJSON)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var tmplFile = "views/main.tmpl"

	tmpl, err := template.New(path.Base(tmplFile)).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, tmpl)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8081", nil)
}
