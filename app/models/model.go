package models

import "encoding/xml"

type TwimlResponse struct {
	XMLName xml.Name `xml:"Response"`
	Gather  *Gather  `xml:",omitempty"`
	Say     *Say     `xml:",omitempty"`
}

type Gather struct {
	XMLName   xml.Name `xml:"Gather"`
	Action    string   `xml:"action,attr"`
	Method    string   `xml:"method,attr"`
	NumDigits int      `xml:"numDigits,attr"`
	Timeout   string   `xml:"timeout,attr,omitempty"`
	Say       *Say     `xml:",omitempty"`
}

type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Text     string   `xml:",chardata"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
}
