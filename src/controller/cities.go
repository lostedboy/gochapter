package controller

import (
	"net/http"
	"../model"
	"../places"
	httpFoundation "../http_foundation"
	"encoding/json"
)

func SuggestionsAction(response http.ResponseWriter, request *http.Request) {
	var citySuggestion model.CitySuggestion

	if request.URL.Query().Get("q") == "" {
		httpFoundation.SendInternalServerErrorResponse(response, "search query is not provided")
		return
	}

	citySuggestion, err := places.FetchAutocomplete(request.URL.Query().Get("q"))

	if err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, err.Error())
		return
	}

	if (citySuggestion.Status != "OK") {
		httpFoundation.SendNotFoundResponse(response)
		return
	}

	httpFoundation.SendJsonResponse(response, "ok", citySuggestion.Predictions)
}

func InfoAction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var placesRequest httpFoundation.PlacesRequest

	err := decoder.Decode(&placesRequest)

	if (placesRequest.Places == nil) || err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, "place_ids is not provided or invalid")
		return
	}

	placeCollection, err := places.FetchPlacesInfo(placesRequest.Places)

	if err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, err.Error())
		return
	}

	httpFoundation.SendJsonResponse(response, "ok", placeCollection.All())
}
