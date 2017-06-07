package model

type CitySuggestion struct {
	Predictions []struct {
		Description string `json:"description"`
		ID string `json:"id"`
		MatchedSubstrings []struct {
			Length int `json:"length"`
			Offset int `json:"offset"`
		} `json:"matched_substrings"`
		PlaceID string `json:"place_id"`
		Reference string `json:"reference"`
		StructuredFormatting struct {
				    MainText string `json:"main_text"`
				    MainTextMatchedSubstrings []struct {
					    Length int `json:"length"`
					    Offset int `json:"offset"`
				    } `json:"main_text_matched_substrings"`
				    SecondaryText string `json:"secondary_text"`
			    } `json:"structured_formatting"`
		Terms []struct {
			Offset int `json:"offset"`
			Value string `json:"value"`
		} `json:"terms"`
		Types []string `json:"types"`
	} `json:"predictions"`
	Status string `json:"status"`
}
