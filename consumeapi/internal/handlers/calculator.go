package handlers

import (
	"bytes"
	"consumeapi/internal/domain/request"
	"consumeapi/internal/domain/response"
	"consumeapi/internal/utils"
	"encoding/xml"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	var numberRequest request.NumberRequest
	err := utils.ReadJSON(r, &numberRequest)
	if err != nil {
		utils.BadRequestResponse(w, err)
		return
	}

	v := utils.NewValidator()
	v.Check(numberRequest.NumberA != 0, "numberA", "numberA is required")
	v.Check(numberRequest.NumberB != 0, "numberB", "numberB is required")

	if !v.Valid() {
		utils.FailedValidationResponse(w, v.Errors)
		return
	}

	addReq := request.AddRequest{Xmlns: "http://tempuri.org/", IntA: numberRequest.NumberA, IntB: numberRequest.NumberB}
	reqBody := request.SoapRequestBody{Request: addReq}
	envelope := request.SoapRequestEnvelope{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Xsi:     "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:     "http://www.w3.org/2001/XMLSchema",
		Body:    reqBody,
	}

	b, err := xml.Marshal(envelope)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer(b))
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", "http://tempuri.org/Add")

	res, err := client.Do(req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	defer res.Body.Close()

	var env response.SoapResponseEnvelope
	err = xml.NewDecoder(res.Body).Decode(&env)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.JSONSuccessResponse{
		Status:  response.Success,
		Message: "two numbers added successfully",
		Data:    env.Body.AddResponse.AddResult,
	})
}
