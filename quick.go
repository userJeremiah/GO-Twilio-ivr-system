package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say string `xml:",omitempty"`
}

func twiml(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{Say: "Welcome, if you can hear this, it works."}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func main() {
	fmt.Println("3....1..liftoff 🚀")
	http.HandleFunc("/twiml", twiml)
	http.ListenAndServe(":3000", nil)
}
