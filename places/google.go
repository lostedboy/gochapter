package places

import (
	"errors"
	"sync"
	"gochapter/config"
	"googlemaps.github.io/maps"
	"context"
)

func GetGoogleClient() (*maps.Client, error) {
	var googleClient *maps.Client

	arguments, err := config.Parse()

	if err != nil {
		return googleClient, errors.New("Google API key not set")
	}

	googleClient, err = maps.NewClient(maps.WithAPIKey(arguments.GoogleKey))

	if err != nil {
		return googleClient, errors.New("Google client can not be created")
	}

	return googleClient, nil
}

func FetchAutocomplete(input string) ([]maps.AutocompletePrediction, error) {
	googleClient, err := GetGoogleClient()

	if err != nil {
		return make([]maps.AutocompletePrediction, 0), err
	}

	autocompleteRequest := maps.PlaceAutocompleteRequest{
		Input : input,
		Language : "en",
		Types : maps.AutocompletePlaceTypeCities,
	}

	result, err := googleClient.PlaceAutocomplete(context.Background(), &autocompleteRequest)

	if err != nil {
		return make([]maps.AutocompletePrediction, 0), err
	}

	return result.Predictions, nil
}

func FetchPlacesInfo(placesIds []string)  ([]maps.PlaceDetailsResult, error) {
	var places []maps.PlaceDetailsResult
	var waitGroup sync.WaitGroup

	googleClient, err := GetGoogleClient()

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

			result, err := googleClient.PlaceDetails(context.Background(), &detailsRequest)

			if err != nil {
				return
			}

			places = append(places, result)
		}(placeId)
	}

	waitGroup.Wait()

	return places, nil
}

func FetchDistanceMatrix(cities []string) ([]maps.DistanceMatrixElementsRow, error) {
	googleClient, err := GetGoogleClient()

	if err != nil {
		return  make([]maps.DistanceMatrixElementsRow, 0) , err
	}

	distanceMatrixRequest := maps.DistanceMatrixRequest{
		Origins  : cities,
		Destinations : cities,
		Language : "en",
	}

	distanceMatrix, _ := googleClient.DistanceMatrix(context.Background(), &distanceMatrixRequest)

	return distanceMatrix.Rows, nil
}

func GetMappedDistanceMatrix(cities []string) (map[string]map[string]maps.Distance, error)  {
	result := make(map[string]map[string]maps.Distance)

	raws, err := FetchDistanceMatrix(cities)

	if err != nil {
		return result, err
	}

	for cityIndex, city := range cities {
		result[city] = make(map[string]maps.Distance)

		for rawIndex, raw := range raws {
			if cityIndex != rawIndex {
				continue
			}

			for elementIndex, element := range raw.Elements {
				if (cityIndex == elementIndex) {
					continue
				}

				result[city][cities[elementIndex]] = element.Distance
			}
		}
	}

	return result, nil
}