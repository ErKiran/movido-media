package main

import (
	"context"
	"fmt"
	"sync"
	"time"

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

	// run at 12 am everyday
	job.AddFunc("0 0 * * *", func() {
		ctx := context.Background()
		// Define your operation function that needs to be retried
		operation := func() error {
			candidates, err := controller.BillingController.SearchCanditates(ctx)
			if err != nil {
				log.Error().Msgf("unable to search billing contriller %s", err)
				return fmt.Errorf("unable to search billing controller: %s", err)
			}

			details, err := controller.BillingController.Details(ctx, candidates)
			if err != nil {
				log.Error().Msgf("unable to get details of contract %s", err)
				return fmt.Errorf("unable to get details of contract: %s", err)
			}

			res, err := controller.BillingController.ExternalBillingCall(ctx)
			if err != nil {
				log.Error().Msgf("mock API request failed %s", err)
				return fmt.Errorf("mock API request failed: %s", err)
			}

			// Mock response utilization
			if res.Type == "debit" {
				for _, det := range details {
					path, err := controller.PDFController.Generate(det)
					if err != nil {
						log.Error().Msgf("unable to generate PDF %s", err)
						return fmt.Errorf("unable to generate PDF: %s", err)
					}

					if err := controller.EmailController.Sender(ctx, det, path); err != nil {
						log.Error().Msgf("unable to send email %s", err)
						return fmt.Errorf("unable to send email: %s", err)
					}
				}
			}

			return nil
		}

		// Call retry function with backoff strategy
		if err := retryWithBackoff(5, time.Second*300, operation); err != nil {
			log.Error().Msgf("operation failed after retries: %s", err)
		}
	})

	var wg sync.WaitGroup
	wg.Add(1)

	job.Start()

	wg.Wait()

	defer utils.Recover()
}

func retryWithBackoff(maxRetries int, initialDelay time.Duration, operation func() error) error {
	delay := initialDelay
	for i := 0; i < maxRetries; i++ {
		if err := operation(); err == nil {
			return nil
		}

		fmt.Printf("Operation failed. Retrying in %s...\n", delay)
		time.Sleep(delay)
		delay *= 2
	}

	return fmt.Errorf("operation failed after %d retries", maxRetries)
}
