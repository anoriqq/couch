package storage

import (
	"context"

	"github.com/anoriqq/couch"
)

type memoryStorage struct {
	v map[string]map[couch.DocID]struct{}
}

var _ couch.Storage = &memoryStorage{}

func (s *memoryStorage) Store(ctx context.Context, docID couch.DocID, keywords []string) error {
	for k, idMap := range s.v {
		for id := range idMap {
			if id == docID {
				delete(s.v[k], docID)
			}
		}
	}

	for _, k := range keywords {
		_, ok := s.v[k]
		if !ok {
			s.v[k] = make(map[string]struct{})
		}

		s.v[k][docID] = struct{}{}
	}

	return nil
}
func (s *memoryStorage) SelectAtOR(ctx context.Context, keywords []string) ([]couch.DocID, error) {
	result := make([]couch.DocID, 0)
	for _, k := range keywords {
		idMap, ok := s.v[k]
		if !ok {
			continue
		}

		ids := make([]string, 0, len(idMap))
		for k2 := range idMap {
			ids = append(ids, k2)
		}

		result = append(result, ids...)
	}

	return result, nil
}

func NewMemory() *memoryStorage {
	return &memoryStorage{
		v: make(map[string]map[couch.DocID]struct{}),
	}
}
