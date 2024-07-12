package main

import (
	"fmt"
	"math/rand"
	"time"

	"pokemon/model" // Import the model package containing the Pokemon struct
)

// Create sample Pokemon
var (
	now = time.Now()
	pokemons = []model.Pokemon{
		{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
		{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 6, Name: "Raymond", Type: "Electric", CatchRate: 80, IsRare: false, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 7, Name: "Didi", Type: "Fire", CatchRate: 50, IsRare: true, RegisteredDate: now.AddDate(0, -5, 0)},
		{ID: 8, Name: "Niko", Type: "Grass/Poison", CatchRate: 70, IsRare: false, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 9, Name: "Bene", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 10, Name: "Aril", Type: "Electric", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -2, 0)},
		{ID: 11, Name: "Agung", Type: "Fire", CatchRate: 60, IsRare: false, RegisteredDate: now.AddDate(0, -7, 0)},
		{ID: 12, Name: "Ahmad", Type: "Grass/Poison", CatchRate: 70, IsRare: false, RegisteredDate: now.AddDate(0, -3, 0)},
		{ID: 13, Name: "Daniel", Type: "Dragon/Flying", CatchRate: 90, IsRare: false, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 14, Name: "Qowi", Type: "Electric", CatchRate: 60, IsRare: false, RegisteredDate: now.AddDate(0, -5, 0)},
	}
)

func main() {

	fmt.Println("Available Pokémon:")
	for _, pokemon := range pokemons {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
			pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
	}

	maxAttempt := 10
	var ids []int
	
	attempt := 0;
	for attempt < maxAttempt {
		var id int
		
		fmt.Print("> Enter Pokémon ID to attempt to catch: ")
		n, err := fmt.Scan(&id)
		if err != nil {
			fmt.Printf("%s, please try again\n\n", err)
			continue
		}

		if n == 1 {
			if !pokemonExists(id, pokemons) {
				fmt.Printf("pokemon with ID %d does not exist, please try again\n\n", id)
				continue
			}
			ids = append(ids, id)
		}

		attempt++
	}

	var selectedPokemons []*model.Pokemon

	for _, id := range ids {
		for _, pokemon := range pokemons {
			if pokemon.ID == id {
				selectedPokemons = append(selectedPokemons, &pokemon)
				break
			}
		}
	}

	for _, selectedPokemon := range selectedPokemons {
		if selectedPokemon != nil {
			result := catchProbability(selectedPokemon.CatchRate)
			fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %d%%: %s\n",
				selectedPokemon.Name, selectedPokemon.Type, selectedPokemon.CatchRate, result)
		} else {
			fmt.Println("Invalid Pokémon ID. Please try again.")
		}
	}
}

func pokemonExists(id int, pokemons []model.Pokemon) bool {
	for _, pokemon := range pokemons {
		if id == pokemon.ID {
			return true
		}
	}
	return false
}

// catchProbability checks if catching is successful given a percentage rate
func catchProbability(rate int) string {
	if rate < 0 || rate > 100 {
		return "Invalid rate. Please provide a rate between 0 and 100."
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS, you caught it"
	}
	return "FAIL, it got away"
}
