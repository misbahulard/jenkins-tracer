package config

import (
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var EsClient *elasticsearch.Client

func ConfigureElasticsearch() {
	var err error

	cfg := elasticsearch.Config{
		Addresses: viper.GetStringSlice("elasticsearch.hosts"),
		Username:  viper.GetString("elasticsearch.username"),
		Password:  viper.GetString("elasticsearch.password"),
	}

	EsClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Errorf("Error when setup elasticsearch client: %s", err)
		os.Exit(1)
	}

	res, err := EsClient.Info()
	if err != nil {
		log.Errorf("Error when conencting elasticsearch cluster: %s", err)
		os.Exit(1)
	}

	if res.StatusCode == 200 {
		log.Info("Elasticsearch: ok")
	}
}
