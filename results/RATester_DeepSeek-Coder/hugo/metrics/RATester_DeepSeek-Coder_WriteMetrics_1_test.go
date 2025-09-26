package metrics

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestWriteMetrics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &Store{
		calculateHints: true,
		metrics: map[string][]time.Duration{
			"metric1": {10 * time.Millisecond, 20 * time.Millisecond, 30 * time.Millisecond},
			"metric2": {40 * time.Millisecond, 50 * time.Millisecond, 60 * time.Millisecond},
		},
		diffs: map[string]*diff{
			"metric1": {baseline: 25 * time.Millisecond, count: 3, simSum: 30},
			"metric2": {baseline: 55 * time.Millisecond, count: 3, simSum: 90},
		},
		cached: map[string]int{
			"metric1": 1,
			"metric2": 2,
		},
	}

	w := &bytes.Buffer{}
	store.WriteMetrics(w)

	got := w.String()
	want := `  cumulative  average  maximum  cache  percent  cached  total 
  ----------  --------  --------  ---------  ------  ------  -----  -------- 
  300ms       100ms     20ms     100       33.33   1      3      metric1
  150ms       50ms      50ms     20        66.67   2      3      metric2
`

	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}
