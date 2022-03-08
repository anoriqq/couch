package couch

import (
	"context"
	"errors"
	"time"

	"github.com/anoriqq/couch/internal/cache"
)

type DocID = string

type CouchSearch struct {
	analyzer             Analyzer
	storage              Storage
	analyzeCacheAtSearch *cache.Cache
}

func (c *CouchSearch) UpdateIndex(ctx context.Context, id DocID, doc string) error {
	if len(id) == 0 {
		return errors.New("id is required")
	}

	keywords, err := c.analyzer.Analyze(ctx, doc)
	if err != nil {
		return err
	}

	err = c.storage.Store(ctx, id, keywords)
	if err != nil {
		return err
	}

	return nil
}

func (c *CouchSearch) Search(ctx context.Context, q string) ([]DocID, error) {
	if len(q) == 0 {
		return nil, errors.New("query string is required")
	}

	keywords, ok := c.analyzeCacheAtSearch.Get(q)
	if !ok {
		var err error
		keywords, err = c.analyzer.Analyze(ctx, q)
		if err != nil {
			return nil, err
		}

		c.analyzeCacheAtSearch.Set(q, keywords)
	}

	result, err := c.storage.SelectAtOR(ctx, keywords)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func New(a Analyzer, s Storage, opts ...option) (*CouchSearch, error) {
	c := new(cfg)

	for _, opt := range opts {
		opt.apply(c)
	}

	return &CouchSearch{
		analyzer:             a,
		storage:              s,
		analyzeCacheAtSearch: cache.New(c.AnalyzerExpirationAtSearch),
	}, nil
}

type cfg struct {
	AnalyzerExpirationAtSearch time.Duration
}

type Analyzer interface {
	Analyze(context.Context, string) ([]string, error)
}

type Storage interface {
	Store(context.Context, DocID, []string) error
	SelectAtOR(context.Context, []string) ([]DocID, error)
}
