package pokeapi

type RespLocations struct{
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct{
		Name string `json:"name"`

	}`json:"results"`
	

}
