package controller

import (
	"net/http"
	"encoding/json"
	"gochapter/places"
	"gochapter/httpFoundation"
)

func SuggestionsAction(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("q") == "" {
		httpFoundation.SendInternalServerErrorResponse(response, "Search query is not provided")
		return
	}

	predictions, err := places.FetchAutocomplete(request.URL.Query().Get("q"))

	if err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, err.Error())
		return
	}

	httpFoundation.SendJsonResponse(response, "ok", predictions)
}

func InfoAction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var placesRequest httpFoundation.PlacesRequest

	err := decoder.Decode(&placesRequest)

	if (placesRequest.Places == nil) || err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, "place_ids is not provided or invalid")
		return
	}

	placesArray, err := places.FetchPlacesInfo(placesRequest.Places)

	if err != nil {
		httpFoundation.SendInternalServerErrorResponse(response, err.Error())
		return
	}

	httpFoundation.SendJsonResponse(response, "ok", placesArray)
}
