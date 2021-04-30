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
	l := log.New(os.Stdout, "git-twitter-bot", log.LstdFlags)

	cfg, err := config.NewBotConfig()
	if err != nil {
		l.Fatal(err)
	}
	oauthCfg := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	oauthToken := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	hc := oauthCfg.Client(oauth1.NoContext, oauthToken)

	tc := twitter.NewClient(hc)

	stopCh := make(chan bool, 1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sigterm is a ^C, handle it
			l.Print("shutting down twitter bot...")
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
				l,
				cfg.TargetOrgUrl,
				func(s string) error {
					tw, res, e := tc.Statuses.Update(s, &twitter.StatusUpdateParams{})
					if e != nil {
						return e
					}
					l.Printf("Response: %s\n Status updated: %v\n", res.Status, tw)
					return nil
				},
			)
			if err != nil {
				l.Print(err)
			}
		}
	}
}
