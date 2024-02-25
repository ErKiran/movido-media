package main

import (
	"movido-media/repositories"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msgf("unable to load environmental file %s", err)
	}

	if err := repositories.Migrate(); err != nil {
		log.Fatal().Msgf("unable to run migration %s", err)
	}
}
