package main

import("fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"

	"github.com/dm1254/pokedexcli/internal/pokeapi"
)


type Pokemon struct{
	Name string

}


type Config struct{
	pokeapiClient pokeapi.Client
	nextPageURL *string 
	previousPageURL *string
	Pokedex map[string]Pokemon
	
}


type clicommand struct{
	name string
	description string
	callback func(*Config, []string) error

}






func getCommand() map[string]clicommand {

	return map[string]clicommand{
		"exit":{
				name : "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
	

		},
		"help":{
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
	


		},
		"map":{
				name: "map",
				description: "Displays locations",
				callback: commandMap,

		},
		"mapb":{
				name: "mapb",
				description: "Displays previous page locations",
				callback: commandMapb,

		
		},
		"explore":{
				name: "explore",
				description: "See pokemon found in current area",
				callback: commandExplore,
	

		},
		"catch":{
				name: "catch",
				description: "Command to catch a pokemon",
				callback: commandCatch,
	
			

		},
		"inspect":{
				name: "inspect",
				description: "Command to inspect pokemon stats",
				callback: commandInspect,


		},
		"pokedex":{
				name: "pokedex",
				description: "Command to view the pokemon you've caught",
				callback: commandPokedex,

				

		},
		}
 
}

func startRepl(cfg *Config){
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}
	for{
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		parts := strings.Fields(text)
		if len(parts) == 0{
			continue

		}
		commandName := strings.ToLower(parts[0])
		args := parts[1:]
		if command, ok := getCommand()[commandName]; ok{
			err := command.callback(config, args)
			if err != nil{
				fmt.Printf("Error executing command: %s\n", err) 
			}
		}else{
			fmt.Println("Unknown command")
		}
		
	}	
}
	

func commandExit(cfg *Config,args []string) error{
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range getCommand(){
		fmt.Printf("%s:%s\n", command.name, command.description)
	}
	return nil
		
}

func commandMap(cfg *Config, args []string) error{
	LocRes,err:= cfg.pokeapiClient.ListLocation(cfg.nextPageURL)
	
	if err != nil{
		return err
	}
	cfg.nextPageURL = LocRes.Next
	cfg.previousPageURL = LocRes.Previous

	for _, Location := range LocRes.Results{
		fmt.Println(Location.Name)
	}
	return nil


}

func commandMapb(cfg *Config, args []string) error{
	if cfg.previousPageURL == nil{
		fmt.Println("You are on the first page!")
	}
	LocRes,err:= cfg.pokeapiClient.ListLocation(cfg.previousPageURL)
	if err != nil{
		return err
	}
	cfg.nextPageURL = LocRes.Next
	cfg.previousPageURL = LocRes.Previous

	for _,Location := range LocRes.Results{
		fmt.Println(Location.Name)
	}
	return nil
}

func commandExplore(cfg *Config, args []string) error{
	areaName := args[0]
	PokRes, err := cfg.pokeapiClient.ListPokemon(areaName)	
	if err != nil{
		return err
	}
	fmt.Printf("Exploring %s...\n",areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range PokRes.PokemonEncounters{
		fmt.Printf("- %s\n",encounter.Pokemon.Name)

	}
	return nil

}


func commandCatch(cfg *Config, args []string) error{
	pokemonName := args[0]
	PokeStat, err := cfg.pokeapiClient.PokemonStats(pokemonName)
	if err != nil{
		return err
	}

	randNum := rand.Intn(700)
	if cfg.Pokedex == nil{
		cfg.Pokedex = make(map[string]Pokemon)	
	}

	fmt.Printf("Throwing a Pokeball at %s...\n",pokemonName)
	if randNum > PokeStat.BaseExperience{
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemonName] = Pokemon{
			Name: pokemonName,
		}

		
	}else{
		fmt.Printf("%s esacped!\n", pokemonName)
	}
	return nil
}

func commandInspect(cfg *Config, args []string) error{
	pokemonName := args[0]
	PokeStat, err := cfg.pokeapiClient.PokemonStats(pokemonName)
	if err != nil{
		return err
	}
	_, exists := cfg.Pokedex[pokemonName]
	if exists{
		fmt.Printf("Name: %s\n", PokeStat.Name)
		fmt.Printf("Height: %d\n",PokeStat.Height)
		fmt.Printf("Weight: %d\n", PokeStat.Weight)
		fmt.Println("Stats:")
		for _, stats := range PokeStat.Stats{
			fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.Base_Stat)	

		}
		fmt.Println("Types:")
		for _, types := range PokeStat.Types{
			fmt.Printf("  -%s\n", types.Type.Name)
		}
	}else{
		fmt.Println("You have not caught that pokemon")
	}
	return nil

}

func commandPokedex(cfg *Config, args []string) error{
	fmt.Println("Your Pokedex")
	for pokemon,_:= range cfg.Pokedex{
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
	

}
