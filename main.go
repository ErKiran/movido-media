package main

import (
	"context"
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

		for _, det := range details {
			path, err := controller.PDFController.Generate(det)
			if err != nil {
				log.Error().Msgf("unable to generate pdf %s", err)
			}

			if err := controller.EmailController.Sender(ctx, det, path); err != nil {
				log.Error().Msgf("unable to send email to  %s", err)
			}
		}
	})

	var wg sync.WaitGroup
	wg.Add(1)

	job.Start()

	wg.Wait()

	defer utils.Recover()
}
