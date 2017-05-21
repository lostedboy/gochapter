package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const googleAppKey = "AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0"

func buildUrl(placeId string) string {
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/place/details/json?key=%s&placeid=%s", googleAppKey, placeId)
}

func main() {
	cityIds := [5]string{
		"ChIJybDUc_xKtUYRTM9XV8zWRD0", // Moscow
		"ChIJAVkDPzdOqEcRcDteW0YgIQQ", // Berlin
		"ChIJdd4hrwug2EcRmSrV3Vo6llI", // London
		"ChIJgTwKgJcpQg0RaSKMYcHeNsQ", // Madrid
		"ChIJw0rXGxGKJRMRAIE4sppPCQM", // Roma
	}

	var cityData [len(cityIds)]string

	for i, v := range cityIds {
		resp, err := http.Get(buildUrl(v))

		if err != nil {
			fmt.Print("Error in fetching response from google")
			continue
		}

		defer resp.Body.Close()

		result, err := ioutil.ReadAll(resp.Body)
		var raw map[string]interface{}
		json.Unmarshal(result, &raw)
		out, _ := json.Marshal(raw)
		cityData[i] = string(out)
	}

	fmt.Printf("%v", cityData)
}