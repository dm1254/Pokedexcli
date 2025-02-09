package pokeapi

import (
	"net/http"	
	"io"
	"encoding/json"
	"time"
	"github.com/dm1254/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocation(pageURL *string) (RespLocations,error){
	url := baseURL + "/location-area"
	cache := pokecache.NewCache(5 * time.Second) 
	if pageURL != nil{
		getCachedData, exists := cache.Get(*pageURL)
		if exists{
			locationResp := RespLocations{}
			if err := json.Unmarshal(getCachedData, &locationResp); err != nil{
				return RespLocations{}, err	
			}
			return locationResp, nil
		}else{
			url = *pageURL
		}
		
		
	}					
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return RespLocations{},err
	}
		
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return RespLocations{}, err
	}

	defer resp.Body.Close()

	data,err := io.ReadAll(resp.Body)
	if err != nil{
		return RespLocations{},err	
	}
	cache.Add(url, data)
	locationResp := RespLocations{}
	if err = json.Unmarshal(data,&locationResp); err != nil{
		return RespLocations{}, err
	}
	return locationResp, nil
}
