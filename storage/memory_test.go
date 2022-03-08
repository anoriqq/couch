package storage_test

import (
	"context"
	"testing"

	"github.com/anoriqq/couch/storage"
	"github.com/google/go-cmp/cmp"
)

func TestMemory(t *testing.T) {
	t.Parallel()

	s := storage.NewMemory()

	docID := "a"
	keywords := []string{"A"}
	err := s.Store(context.Background(), docID, keywords)

	if err != nil {
		t.Errorf("err := s.Store(context.Background(), docID, keywords); err = %v want nil", err)
	}

	result, err := s.SelectAtOR(context.Background(), keywords)

	if err != nil {
		t.Errorf("result, err := s.SelectAtOR(context.Background(), keywords); err = %v want nil", err)
	}
	if diff := cmp.Diff([]string{docID}, result); diff != "" {
		t.Errorf("result, err := s.SelectAtOR(context.Background(), keywords); result mismatch  (-want +got):\n%s", diff)
	}
}
