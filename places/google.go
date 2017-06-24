package places

import (
	"errors"
	"sync"
	"gochapter/config"
	"googlemaps.github.io/maps"
	"context"
)

func GetGoogleClient() (*maps.Client, error) {
	var googleCLient *maps.Client

	arguments, err := config.Parse()

	if err != nil {
		return googleCLient, errors.New("Google API key not set")
	}

	googleCLient, err = maps.NewClient(maps.WithAPIKey(arguments.GoogleKey))

	if err != nil {
		return googleCLient, errors.New("Google client can not be created")
	}

	return googleCLient, nil
}

func FetchAutocomplete(input string) ([]maps.AutocompletePrediction, error) {
	googleCLient, err := GetGoogleClient()

	if err != nil {
		return make([]maps.AutocompletePrediction, 0), err
	}

	autocompleteRequest := maps.PlaceAutocompleteRequest{
		Input : input,
		Language : "en",
		Types : maps.AutocompletePlaceTypeCities,
	}

	result, err := googleCLient.PlaceAutocomplete(context.Background(), &autocompleteRequest)

	if err != nil {
		return make([]maps.AutocompletePrediction, 0), err
	}

	return result.Predictions, nil
}

func FetchPlacesInfo(placesIds []string)  ([]maps.PlaceDetailsResult, error) {
	var places []maps.PlaceDetailsResult
	var waitGroup sync.WaitGroup

	googleCLient, err := GetGoogleClient()

	if err != nil {
		return places, err
	}

	for _, placeId := range placesIds {
		waitGroup.Add(1)

		go func(placeId string) {
			defer waitGroup.Done()

			detailsRequest := maps.PlaceDetailsRequest{
				PlaceID  : placeId,
				Language : "en",
			}

			result, err := googleCLient.PlaceDetails(context.Background(), &detailsRequest)

			if err != nil {
				return
			}

			places = append(places, result)
		}(placeId)
	}

	waitGroup.Wait()

	return places, nil
}

func FetchDistaceMatrix(origins []string)  {
	
}