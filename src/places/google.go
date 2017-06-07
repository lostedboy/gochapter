package places

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../model"
	"errors"
)

const autocompleteUrl = "https://maps.googleapis.com/maps/api/place/autocomplete/json?key=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0&types=(cities)&language=en&input=%s"

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
