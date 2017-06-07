package response

import (
	"net/http"
	"encoding/json"
)

type NotFoundResponse struct {
	Status    string `json:"status"`
}

type InternalServerErrorResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type JsonResponse struct {
	Status    string `json:"status"`
	Data      interface{} `json:"data"`
}

func SendNotFoundResponse(response http.ResponseWriter)  {
	notFoundResponse := NotFoundResponse{"not_found"}
	jsonString, _ := json.Marshal(notFoundResponse)

	SendResponse(response, 404, jsonString)
}

func SendInternalServerErrorResponse(response http.ResponseWriter, message string)  {
	internalServerErrorResponse := InternalServerErrorResponse{"error", message}
	jsonString, _ := json.Marshal(internalServerErrorResponse)

	SendResponse(response, 500, jsonString)
}

func SendJsonResponse(response http.ResponseWriter, status string, data interface{})  {
	jsonResponse := JsonResponse{status, data}

	jsonString, _ := json.Marshal(jsonResponse)

	SendResponse(response, 200, jsonString)
}

func SendResponse(response http.ResponseWriter, code int, jsonString []byte)  {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(jsonString)
}
