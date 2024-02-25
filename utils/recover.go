package utils

import (
	"github.com/rs/zerolog/log"
)

// Recover recovers from panics and logs the error.
func Recover() {
	if r := recover(); r != nil {
		log.Info().Msgf("Recovered from the panic %s", r)
	}
}
