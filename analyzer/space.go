package analyzer

import (
	"context"
	"strings"

	"github.com/anoriqq/couch"
)

type spaceAnalyzer struct{}

var _ couch.Analyzer = spaceAnalyzer{}

func (spaceAnalyzer) Analyze(ctx context.Context, v string) ([]string, error) {
	keywords := strings.Split(v, " ")

	return keywords, nil
}

func NewSpace() *spaceAnalyzer {
	return &spaceAnalyzer{}
}
