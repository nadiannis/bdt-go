package response

import "encoding/xml"

type ResponseStatus string

var (
	Success ResponseStatus = "success"
	Error   ResponseStatus = "error"
)

type JSONErrorResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Detail  any            `json:"detail"`
}

type JSONSuccessResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    any            `json:"data"`
}

type SoapResponseEnvelope struct {
	XMLName xml.Name         `xml:"Envelope"`
	SoapEnv string           `xml:"xmlns:soap,attr"`
	Xsi     string           `xml:"xmlns:xsi,attr"`
	Xsd     string           `xml:"xmlns:xsd,attr"`
	Body    SoapResponseBody `xml:"Body"`
}

type SoapResponseBody struct {
	XMLName     xml.Name    `xml:"Body"`
	AddResponse AddResponse `xml:"AddResponse"`
}

type AddResponse struct {
	XMLName   xml.Name `xml:"AddResponse"`
	AddResult int      `xml:"AddResult"`
}
