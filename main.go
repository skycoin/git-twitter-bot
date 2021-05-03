package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Skycoin/git-telegram-bot/pkg/githandler"
	"github.com/Skycoin/git-twitter-bot/internal/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	logger := log.New(os.Stdout, "git-twitter-bot", log.LstdFlags)

	cfg, err := config.NewBotConfig()
	if err != nil {
		logger.Fatal(err)
	}
	oauthCfg := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	oauthToken := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	httpClient := oauthCfg.Client(oauth1.NoContext, oauthToken)

	twitterClient := twitter.NewClient(httpClient)

	stopCh := make(chan bool, 1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sigterm is a ^C, handle it
			logger.Print("shutting down twitter bot...")
			time.Sleep(3 * time.Second)
			stopCh <- true
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	var previousEventId string
	var currentEventId string

	for {
		select {
		case <-stopCh:
			ticker.Stop()
			break
		case <-ticker.C:
			err = githandler.HandleStartCommand(
				previousEventId,
				currentEventId,
				logger,
				cfg.TargetOrgUrl,
				func(s string) error {
					tweet, res, tweetErr := twitterClient.Statuses.Update(s, &twitter.StatusUpdateParams{})
					if tweetErr != nil {
						return tweetErr
					}
					logger.Printf("Response: %s\n Status updated: %v\n", res.Status, tweet)
					return nil
				},
			)
			if err != nil {
				logger.Print(err)
			}
		}
	}
}
