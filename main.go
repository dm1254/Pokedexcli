package main 
import(
	"github.com/dm1254/pokedexcli/internal/pokeapi"
	"time"
	"math/rand"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &Config{
		pokeapiClient: pokeClient,

	}
	startRepl(cfg)	
}

<<<<<<< HEAD
=======

>>>>>>> a9024b1 (Update main.go)
