package cache_test

import (
	"testing"
	"time"

	"github.com/anoriqq/couch/internal/cache"
	"github.com/google/go-cmp/cmp"
	"github.com/tenntenn/testtime"
)

func TestCache(t *testing.T) {
	t.Parallel()

	c := cache.New(time.Minute)

	testtime.SetTime(t, parseTime(t, "2022-02-02T02:02:02Z"))

	key := "key"
	value := []string{"A"}
	c.Set(key, value)

	got, ok := c.Get(key)

	if !ok {
		t.Errorf("Get(\"a\") = %t want %t", ok, true)
	}

	if diff := cmp.Diff(value, got); diff != "" {
		t.Errorf("Get(\"a\") mismatch (-want +got):\n%s", diff)
	}

	testtime.SetTime(t, parseTime(t, "2022-02-02T02:02:02Z").Add(time.Hour))

	got, ok = c.Get(key)

	if ok {
		t.Errorf("Get(\"a\") = %t want %t", ok, false)
	}

	if diff := cmp.Diff([]string(nil), got); diff != "" {
		t.Errorf("Get(\"a\") mismatch (-want +got):\n%s", diff)
	}
}

func parseTime(t *testing.T, v string) time.Time {
	t.Helper()

	tm, err := time.Parse("2006-01-02T15:04:05Z", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return tm
}
