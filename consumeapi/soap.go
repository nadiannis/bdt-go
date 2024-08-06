package main

import (
	"bytes"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NumberToWordsResponse struct {
	XMLName xml.Name `xml:"NumberToWordsResponse"`
	NumberToWordsResult string `xml:"NumberToWordsResult"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"Body"`
	NumberToWordsResponse NumberToWordsResponse `xml:"NumberToWordsResponse"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body SOAPBody `xml:"Body"`
}

func consumeSOAP(c *gin.Context) {
	soapRequest := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
				<ubiNum>520</ubiNum>
			</NumberToWords>
		</soap:Body>
	</soap:Envelope>`

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.dataaccess.com/webservicesserver/NumberConversion.wso", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to make request"})
		return
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response"})
	// 	return
	// }

	// c.Data(http.StatusOK, "text/xml", body)

	var envelope SOAPEnvelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode response", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"NumberToWordsResult": envelope.Body.NumberToWordsResponse.NumberToWordsResult,
	})
}
