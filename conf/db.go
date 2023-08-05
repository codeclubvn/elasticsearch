package conf

import (
	"github.com/cenkalti/backoff"
	es "github.com/elastic/go-elasticsearch/v7"
	"log"
	"time"
)

func (a *App) GetDB() (esClient *es.Client, err error) {
	retryBackoff := backoff.NewExponentialBackOff()
	cfg := es.Config{
		// Retry on 429 TooManyRequests statuses
		RetryOnStatus: []int{502, 503, 504, 429},
		// Configure the backoff function
		//
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		// Retry up to 5 attempts
		//
		MaxRetries: 5,
		Addresses:  []string{GetEnv().ESDomain},
		//Username:   utils.ESUsername,
		//Password:   utils.ESPassword,
	}
	clientES, err := es.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(clientES.Info())
	}

	return clientES, err
}
