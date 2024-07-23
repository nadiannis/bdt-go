package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	log.Print("Hello, world!")

	// log.Panic().Msg("Hello, world!")

	// log.Fatal().Msg("Hello, world!")

	log.Error().Msg("Hello, world!")

	log.Warn().Msg("Hello, world!")

	log.Info().Msg("Hello, world!")

	log.Debug().Msg("Hello, world!")

	log.Trace().Msg("Hello, world!")

	log.Debug().
		Str("name", "Nadia").
		Send()

	log.Debug().
		Str("name", "Nadia").
		Int("age", 24).
		Msg("hi")
}
