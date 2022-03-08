package analyzer_test

import (
	"context"
	"testing"

	"github.com/anoriqq/couch/analyzer"
	"github.com/google/go-cmp/cmp"
)

func TestSpace(t *testing.T) {
	t.Parallel()

	a := analyzer.NewSpace()

	tests := map[string]struct {
		in  string
		out []string
	}{
		"1 words": {
			in:  "a",
			out: []string{"a"},
		},
		"2 words": {
			in:  "a b",
			out: []string{"a", "b"},
		},
	}
	for k, tt := range tests {
		tt := tt
		t.Run(k, func(t *testing.T) {
			t.Parallel()

			got, err := a.Analyze(context.Background(), tt.in)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tt.out, got); diff != "" {
				t.Errorf("got, err := a.Analyze(context.Background(), %s); got = %v want %v", tt.in, got, tt.out)
			}
		})
	}
}
