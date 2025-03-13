package elas

import (
	"app/config"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchClient struct {
	Client    *elasticsearch7.Client
	IndexName string
}

func NewElasticsearchClient(cfg *config.Config) (*ElasticsearchClient, error) {
	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: cfg.Elasticsearch.Addresses,
	})
	if err != nil {
		return nil, err
	}
	return &ElasticsearchClient{
		Client:    client,
		IndexName: cfg.Elasticsearch.IndexName,
	}, nil
}
