package model

type Place struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result struct {
				 AddressComponents []struct {
					 LongName string `json:"long_name"`
					 ShortName string `json:"short_name"`
					 Types []string `json:"types"`
				 } `json:"address_components"`
				 AdrAddress string `json:"adr_address"`
				 FormattedAddress string `json:"formatted_address"`
				 Geometry struct {
							   Location struct {
									    Lat float64 `json:"lat"`
									    Lng float64 `json:"lng"`
								    } `json:"location"`
							   Viewport struct {
									    Northeast struct {
											      Lat float64 `json:"lat"`
											      Lng float64 `json:"lng"`
										      } `json:"northeast"`
									    Southwest struct {
											      Lat float64 `json:"lat"`
											      Lng float64 `json:"lng"`
										      } `json:"southwest"`
								    } `json:"viewport"`
						   } `json:"geometry"`
				 Icon string `json:"icon"`
				 ID string `json:"id"`
				 Name string `json:"name"`
				 Photos []struct {
					 Height int `json:"height"`
					 HTMLAttributions []string `json:"html_attributions"`
					 PhotoReference string `json:"photo_reference"`
					 Width int `json:"width"`
				 } `json:"photos"`
				 PlaceID string `json:"place_id"`
				 Reference string `json:"reference"`
				 Scope string `json:"scope"`
				 Types []string `json:"types"`
				 URL string `json:"url"`
				 UtcOffset int `json:"utc_offset"`
				 Vicinity string `json:"vicinity"`
			 } `json:"result"`
	Status string `json:"status"`
}

type PlaceCollection struct {
	Items []Place

}

func (placeCollection *PlaceCollection) Add(place Place) {
	placeCollection.Items = append(placeCollection.Items, place)
}

func (placeCollection *PlaceCollection) All() []Place {
	return placeCollection.Items
}