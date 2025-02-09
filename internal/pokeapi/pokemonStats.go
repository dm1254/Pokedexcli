package pokeapi 

import(
	"net/http"	
	"io"
	"encoding/json"
	"time"
	"github.com/dm1254/pokedexcli/internal/pokecache"

)

func(c *Client) PokemonStats(pokemonName string) (pokeStatResp, error){
	url := baseURL + "/pokemon/" + pokemonName
	cache := pokecache.NewCache(5 * time.Second)
	data, exists := cache.Get(url)
	if exists{
		pokemonStatresp := pokeStatResp{}
		if err := json.Unmarshal(data, &pokemonStatresp); err != nil{
			return pokeStatResp{}, err
		}
		return pokemonStatresp, nil
		
	}
	req, err := http.NewRequest("GET",url, nil)
	if err != nil{
		return pokeStatResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return pokeStatResp{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return pokeStatResp{}, err
	}
	cache.Add(url,data)
	pokemonStatresp := pokeStatResp{}
	if err = json.Unmarshal(data, &pokemonStatresp); err != nil{
		return pokeStatResp{},err
	}
	return pokemonStatresp, nil

	
}

