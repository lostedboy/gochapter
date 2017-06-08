package places

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../model"
	"errors"
	"sync"
)

const autocompleteUrl = "https://maps.googleapis.com/maps/api/place/autocomplete/json?key=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0&types=(cities)&language=en&input=%s"
const placesUrl  = "https://maps.googleapis.com/maps/api/place/details/json?key=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0&placeid=%s"

func FetchAutocomplete(input string) (model.CitySuggestion, error) {
	var citySuggestion model.CitySuggestion

	resp, err := http.Get(fmt.Sprintf(autocompleteUrl, input))

	if err != nil {
		return citySuggestion, errors.New("Error in fetching response from google")
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return citySuggestion, errors.New("Error in reading response from google")
	}

	err = json.Unmarshal(result, &citySuggestion)

	if err != nil {
		return citySuggestion, errors.New("Error in parsing google response")
	}

	return citySuggestion, nil
}

func FetchPlacesInfo(placesIds []string)  (*model.PlaceCollection, error) {
	var places = new(model.PlaceCollection)
	var waitGroup sync.WaitGroup

	for _, placeId := range placesIds {
		waitGroup.Add(1)

		go func(placeId string) {
			defer waitGroup.Done()

			resp, err := http.Get(fmt.Sprintf(placesUrl, placeId))

			if err != nil {
				return
			}

			defer resp.Body.Close()

			result, err := ioutil.ReadAll(resp.Body)

			var place model.Place

			json.Unmarshal(result, &place)

			places.Add(place)
		}(placeId)
	}

	waitGroup.Wait()

	return places, nil
}
