package controller

import (
	"net/http"
	"../model"
	"../places"
	httpResponse "../response"
)

func SuggestionsAction(response http.ResponseWriter, request *http.Request) {
	var citySuggestion model.CitySuggestion

	citySuggestion, err := places.FetchAutocomplete(request.URL.Query().Get("q"))

	if err != nil {
		httpResponse.SendInternalServerErrorResponse(response, err.Error())
		return
	}

	if (citySuggestion.Status != "OK") {
		httpResponse.SendNotFoundResponse(response)
		return
	}

	httpResponse.SendJsonResponse(response, "ok", citySuggestion.Predictions)
}
