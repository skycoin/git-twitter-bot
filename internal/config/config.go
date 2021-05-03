package config

import (
	"net/url"
	"os"

	"github.com/Skycoin/git-telegram-bot/pkg/errutil"
)

const (
	defaultEventCount = "3"
)

type BotConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	TargetOrgUrl   string
}

func NewBotConfig() (*BotConfig, error) {
	consumerKey := os.Getenv("TW_CONSUMER_KEY")
	if consumerKey == "" {
		return nil, errutil.ErrInvalidConfig.Desc("empty consumer key")
	}
	consumerSecret := os.Getenv("TW_CONSUMER_SECRET")
	if consumerSecret == "" {
		return nil, errutil.ErrInvalidConfig.Desc("empty consumer secret")
	}
	accessToken := os.Getenv("TW_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, errutil.ErrInvalidConfig.Desc("empty access token")
	}
	accessTokenSecret := os.Getenv("TW_ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		return nil, errutil.ErrInvalidConfig.Desc("empty access token secret")
	}
	targetOrgUrl := os.Getenv("TW_TARGET_ORG_URL")
	if targetOrgUrl == "" {
		return nil, errutil.ErrInvalidConfig.Desc("empty target org url")
	}
	bc := &BotConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessTokenSecret,
		TargetOrgUrl:   targetOrgUrl,
	}

	u, err := url.Parse(bc.TargetOrgUrl)
	if err != nil {
		return nil, errutil.ErrInvalidUrl.Desc(bc.TargetOrgUrl)
	}
	// set url param for per_page item
	q := u.Query()
	q.Set("per_page", defaultEventCount)
	u.RawQuery = q.Encode()
	bc.TargetOrgUrl = u.String()
	return bc, nil
}
