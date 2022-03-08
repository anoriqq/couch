package couch

import "time"

type option interface {
	apply(*cfg)
}

func WithAnalyzerExpirationAtSearch(e time.Duration) withAnalyzerExpirationAtSearch {
	return withAnalyzerExpirationAtSearch{expiration: e}
}

type withAnalyzerExpirationAtSearch struct {
	expiration time.Duration
}

func (e withAnalyzerExpirationAtSearch) apply(c *cfg) {
	c.AnalyzerExpirationAtSearch = e.expiration
}
