package pokeapi
type pokeStatResp struct{
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct{
		Base_Stat int `json:"base_stat"` 
		Stat struct{
			Name string `json:"name"`
		}`json:"stat"`
	}`json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`

}

