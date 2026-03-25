package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/gaokao-advisor/backend/internal/config"
	"github.com/gaokao-advisor/backend/internal/model"
)

type RAGService struct {
	esClient *elasticsearch.Client
}

func NewRAGService(esClient *elasticsearch.Client) *RAGService {
	return &RAGService{esClient: esClient}
}

type CollegeDoc struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Type        string    `json:"type"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	Majors      []string  `json:"majors"`
	Rankings    string    `json:"rankings"`
	Embedding   []float32 `json:"embedding,omitempty"`
}

type MajorDoc struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Code        string    `json:"code"`
	Degree      string    `json:"degree"`
	Description string    `json:"description"`
	CoreCourses []string  `json:"core_courses"`
	Employment  string    `json:"employment"`
	Embedding   []float32 `json:"embedding,omitempty"`
}

type KnowledgeDoc struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	Embedding []float32 `json:"embedding,omitempty"`
}

type SearchResult struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Score   float64 `json:"score"`
}

func (s *RAGService) CreateCollegeIndex(ctx context.Context) error {
	mapping := `{
		"settings": {
			"analysis": {
				"analyzer": {
					"ik_analyzer": {
						"type": "standard",
						"stopwords": "_none_"
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": {"type": "keyword"},
				"name": {"type": "text", "analyzer": "ik_analyzer"},
				"province": {"type": "keyword"},
				"city": {"type": "keyword"},
				"type": {"type": "keyword"},
				"level": {"type": "keyword"},
				"description": {"type": "text", "analyzer": "ik_analyzer"},
				"majors": {"type": "keyword"},
				"rankings": {"type": "text"},
				"embedding": {"type": "dense_vector", "dims": 1536, "index": true, "similarity": "cosine"}
			}
		}
	}`

	res, err := s.esClient.Indices.Create(
		"colleges",
		s.esClient.Indices.Create.WithContext(ctx),
		s.esClient.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}
	return nil
}

func (s *RAGService) CreateMajorIndex(ctx context.Context) error {
	mapping := `{
		"settings": {
			"analysis": {
				"analyzer": {
					"ik_analyzer": {
						"type": "standard"
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": {"type": "keyword"},
				"name": {"type": "text", "analyzer": "ik_analyzer"},
				"category": {"type": "keyword"},
				"code": {"type": "keyword"},
				"degree": {"type": "keyword"},
				"description": {"type": "text", "analyzer": "ik_analyzer"},
				"core_courses": {"type": "keyword"},
				"employment": {"type": "text"},
				"embedding": {"type": "dense_vector", "dims": 1536, "index": true, "similarity": "cosine"}
			}
		}
	}`

	res, err := s.esClient.Indices.Create(
		"majors",
		s.esClient.Indices.Create.WithContext(ctx),
		s.esClient.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (s *RAGService) CreateKnowledgeIndex(ctx context.Context) error {
	mapping := `{
		"settings": {
			"analysis": {
				"analyzer": {
					"ik_analyzer": {
						"type": "standard"
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"id": {"type": "keyword"},
				"type": {"type": "keyword"},
				"title": {"type": "text", "analyzer": "ik_analyzer"},
				"content": {"type": "text", "analyzer": "ik_analyzer"},
				"tags": {"type": "keyword"},
				"embedding": {"type": "dense_vector", "dims": 1536, "index": true, "similarity": "cosine"}
			}
		}
	}`

	res, err := s.esClient.Indices.Create(
		"knowledge_base",
		s.esClient.Indices.Create.WithContext(ctx),
		s.esClient.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (s *RAGService) IndexCollege(ctx context.Context, college *model.College) error {
	doc := CollegeDoc{
		ID:          fmt.Sprintf("%d", college.ID),
		Name:        college.Name,
		Province:    college.Province,
		City:        college.City,
		Type:        college.Type,
		Level:       college.Level,
		Description: college.Description,
		Majors:      strings.Split(college.Disciplines, ","),
		Rankings:    college.Rankings,
	}

	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	res, err := s.esClient.Index(
		"colleges",
		bytes.NewReader(data),
		s.esClient.Index.WithContext(ctx),
		s.esClient.Index.WithDocumentID(fmt.Sprintf("%d", college.ID)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (s *RAGService) IndexMajor(ctx context.Context, major *model.Major) error {
	doc := MajorDoc{
		ID:          fmt.Sprintf("%d", major.ID),
		Name:        major.Name,
		Category:    major.Category,
		Code:        major.Code,
		Degree:      major.Degree,
		Description: major.Description,
		CoreCourses: strings.Split(major.CoreCourses, ","),
		Employment:  major.Employment,
	}

	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	res, err := s.esClient.Index(
		"majors",
		bytes.NewReader(data),
		s.esClient.Index.WithContext(ctx),
		s.esClient.Index.WithDocumentID(fmt.Sprintf("%d", major.ID)),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (s *RAGService) Search(ctx context.Context, query string, topK int) ([]SearchResult, error) {
	searchQuery := fmt.Sprintf(`{
		"query": {
			"multi_match": {
				"query": %q,
				"fields": ["name^2", "description", "title", "content"],
				"type": "best_fields"
			}
		},
		"size": %d
	}`, query, topK)

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(ctx),
		s.esClient.Search.WithIndex("colleges", "majors", "knowledge_base"),
		s.esClient.Search.WithBody(strings.NewReader(searchQuery)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	var results []SearchResult
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		h := hit.(map[string]interface{})
		source := h["_source"].(map[string]interface{})
		results = append(results, SearchResult{
			ID:      h["_id"].(string),
			Type:    h["_index"].(string),
			Title:   s.getStringField(source, "name", "title"),
			Content: s.getStringField(source, "description", "content"),
			Score:   h["_score"].(float64),
		})
	}

	return results, nil
}

func (s *RAGService) getStringField(doc map[string]interface{}, fields ...string) string {
	for _, field := range fields {
		if val, ok := doc[field]; ok {
			if str, ok := val.(string); ok {
				return str
			}
		}
	}
	return ""
}

func (s *RAGService) HybridSearch(ctx context.Context, query string, topK int) ([]SearchResult, error) {
	bm25Query := fmt.Sprintf(`{
		"query": {
			"bool": {
				"should": [
					{
						"multi_match": {
							"query": %q,
							"fields": ["name^2", "description", "title^1.5", "content"],
							"type": "best_fields"
						}
					}
				]
			}
		},
		"size": %d
	}`, query, topK)

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(ctx),
		s.esClient.Search.WithIndex("colleges", "majors", "knowledge_base"),
		s.esClient.Search.WithBody(strings.NewReader(bm25Query)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	var results []SearchResult
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		h := hit.(map[string]interface{})
		source := h["_source"].(map[string]interface{})
		results = append(results, SearchResult{
			ID:      h["_id"].(string),
			Type:    h["_index"].(string),
			Title:   s.getStringField(source, "name", "title"),
			Content: s.getStringField(source, "description", "content"),
			Score:   h["_score"].(float64),
		})
	}

	return results, nil
}

func NewESClient(cfg *config.Config) (*elasticsearch.Client, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{cfg.ES.Hosts},
	}

	if cfg.ES.Username != "" {
		esCfg.Username = cfg.ES.Username
		esCfg.Password = cfg.ES.Password
	}

	return elasticsearch.NewClient(esCfg)
}
