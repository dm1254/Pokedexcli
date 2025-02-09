package pokeapi 

import(
	"net/http"	
	"io"
	"encoding/json"
	"time"
	"github.com/dm1254/pokedexcli/internal/pokecache"

)

func(c *Client) ListPokemon(areaURL string) (RespPokemon, error){
	url := baseURL + "/location-area/" + areaURL
	cache := pokecache.NewCache(5 * time.Second)
	data, exists := cache.Get(url)
	if exists{
		pokemonresp := RespPokemon{}
		if err := json.Unmarshal(data, &pokemonresp); err != nil{
			return RespPokemon{}, err
		}
		return pokemonresp, nil
		
	}
	req, err := http.NewRequest("GET",url, nil)
	if err != nil{
		return RespPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return RespPokemon{}, err
	}
	cache.Add(url,data)
	pokemonresp := RespPokemon{}
	if err = json.Unmarshal(data, &pokemonresp); err != nil{
		return RespPokemon{},err
	}
	return pokemonresp, nil

	
}

