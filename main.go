package main

import (
	"context"
	"fmt"
	"sync"

	"movido-media/controllers"
	"movido-media/repositories"
	"movido-media/utils"

	"movido-media/cron"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// setting global log level to info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if err := godotenv.Load(); err != nil {
		log.Error().Msgf("unable to get environment variables %s", err)
	}

	db, err := repositories.Connect()
	if err != nil {
		log.Error().Msgf("unable to connect to db %s", err)
	}

	controller := controllers.NewController(db)

	job := cron.NewCronJob()

	job.AddFunc("@every 10s", func() {
		ctx := context.Background()
		candidates, err := controller.BillingController.SearchCanditates(ctx)
		if err != nil {
			log.Error().Msgf("unable to search billing controller %s", err)
		}

		details, err := controller.BillingController.Details(ctx, candidates)
		if err != nil {
			log.Error().Msgf("unable to get details of contract %s", err)
		}

		fmt.Println("Halo und Tschuss!!", details)
	})

	var wg sync.WaitGroup
	wg.Add(1)

	job.Start()

	wg.Wait()

	defer utils.Recover()
}
