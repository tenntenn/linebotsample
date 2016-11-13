package linebot

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type ChannelConfig struct {
	Secret string `datastore:"Secret"`
	Token  string `datastore:"Token"`
}

func LoadConfig(c context.Context) (*ChannelConfig, error) {
	var config []*ChannelConfig
	_, err := datastore.NewQuery("ChannelConfig").Limit(1).GetAll(c, &config)
	if err != nil {
		return nil, err
	}

	if len(config) < 1 {
		return nil, errors.New("cannnot load channel config")
	}

	return config[0], nil
}
