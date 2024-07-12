package main

import (
	"fmt"
	"math/rand"
	"time"

	"pokemon/model" // Import the model package containing the Pokemon struct
)

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

func main() {
	now := time.Now()

	// Create sample Pokemon
	pokemons := []model.Pokemon{
		{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
		{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
	}

	fmt.Println("Available Pokémon:")
	for _, pokemon := range pokemons {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
			pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
	}

	var id int
	fmt.Print("Enter Pokémon ID to attempt to catch: ")
	fmt.Scan(&id)

	var selectedPokemon *model.Pokemon
	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			selectedPokemon = &pokemon
			break
		}
	}

	if selectedPokemon != nil {
		result := catchProbability(selectedPokemon.CatchRate)
		fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %d%%: %s\n",
			selectedPokemon.Name, selectedPokemon.Type, selectedPokemon.CatchRate, result)
	} else {
		fmt.Println("Invalid Pokémon ID. Please try again.")
	}
}
