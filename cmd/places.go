package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
	"../model"
)

const urlTemplate  = "https://maps.googleapis.com/maps/api/place/details/json?key=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0&placeid=%s"

func main() {
	cityIds := [5]string{
		"ChIJybDUc_xKtUYRTM9XV8zWRD0", // Moscow
		"ChIJAVkDPzdOqEcRcDteW0YgIQQ", // Berlin
		"ChIJdd4hrwug2EcRmSrV3Vo6llI", // London
		"ChIJgTwKgJcpQg0RaSKMYcHeNsQ", // Madrid
		"ChIJw0rXGxGKJRMRAIE4sppPCQM", // Roma
	}

	var places = new(model.PlaceCollection)
	var waitGroup sync.WaitGroup

	for _, placeId := range cityIds {
		waitGroup.Add(1)

		go func(placeId string) {
			defer waitGroup.Done()

			resp, err := http.Get(fmt.Sprintf(urlTemplate, placeId))

			if err != nil {
				fmt.Print("Error in fetching response from google")

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

	placesJson, _ := json.Marshal(places.All())

	ioutil.WriteFile("output.json", placesJson, 0644)

	fmt.Println("Done")
}