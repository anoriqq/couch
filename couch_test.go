package couch_test

import (
	"context"
	"testing"
	"time"

	"github.com/anoriqq/couch"
	"github.com/anoriqq/couch/analyzer"
	"github.com/anoriqq/couch/storage"
	"github.com/google/go-cmp/cmp"
)

func TestCouch(t *testing.T) {
	t.Parallel()

	c, err := couch.New(analyzer.NewSpace(), storage.NewMemory(), couch.WithAnalyzerExpirationAtSearch(10*time.Second))

	if err != nil {
		t.Error("unexpected error:", err)
	}

	id1 := "a+b"
	doc := "A B"
	err = c.UpdateIndex(context.Background(), id1, doc)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	q := "A"
	result, err := c.Search(context.Background(), q)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	want := []string{id1}
	if diff := cmp.Diff(want, result); diff != "" {
		t.Fatalf("result, err := c.Search(context.Background(), %s); result = %v want %v", q, result, want)
	}

	id2 := "a+z"
	doc = "A Z"
	err = c.UpdateIndex(context.Background(), id2, doc)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	q = "A"
	result, err = c.Search(context.Background(), q)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	want = []string{id1, id2}
	if diff := cmp.Diff(want, result); diff != "" {
		t.Fatalf("result, err := c.Search(context.Background(), %s); result = %v want %v", q, result, want)
	}
}
