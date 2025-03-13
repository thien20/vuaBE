package repository

import (
	"app/internal/infra/elas"
	"bytes"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type SearchRepositoryInterface interface {
	SearchSimple(keyword string) ([]string, error)
	// SearchSemantic(keyword string) ([]string, error)
}

type SearchRepository struct {
	db   *gorm.DB
	elas *elas.ElasticsearchClient
}

func NewSearchRepository(db *gorm.DB, elas *elas.ElasticsearchClient) *SearchRepository {
	return &SearchRepository{
		db:   db,
		elas: elas}
}

func (r *SearchRepository) SearchSimple(keyword string) ([]string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": keyword,
			},
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := r.elas.Client.Search(
		r.elas.Client.Search.WithIndex(r.elas.IndexName), // thinking more about this
		r.elas.Client.Search.WithBody(&buf),
		r.elas.Client.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	hitsData, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected data type for hits")
	}

	contents := make([]string, 0, len(hitsData))
	for _, hit := range hitsData {
		src, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if ok {
			contents = append(contents, fmt.Sprintf("%v", src["content"]))
		}
	}
	return contents, nil
}
