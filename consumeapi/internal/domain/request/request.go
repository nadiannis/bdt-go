package request

import "encoding/xml"

type NumberRequest struct {
	NumberA int `json:"numberA"`
	NumberB int `json:"numberB"`
}

type SoapRequestEnvelope struct {
	XMLName xml.Name        `xml:"soap:Envelope"`
	SoapEnv string          `xml:"xmlns:soap,attr"`
	Xsi     string          `xml:"xmlns:xsi,attr"`
	Xsd     string          `xml:"xmlns:xsd,attr"`
	Body    SoapRequestBody `xml:"soap:Body"`
}

type SoapRequestBody struct {
	XMLName xml.Name   `xml:"soap:Body"`
	Request AddRequest `xml:"Add"`
}

type AddRequest struct {
	XMLName xml.Name `xml:"Add"`
	Xmlns   string   `xml:"xmlns,attr"`
	IntA    int      `xml:"intA"`
	IntB    int      `xml:"intB"`
}
