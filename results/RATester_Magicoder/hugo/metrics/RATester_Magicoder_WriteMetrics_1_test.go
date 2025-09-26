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
			"metric1": {time.Second, time.Second * 2, time.Second * 3},
			"metric2": {time.Second * 4, time.Second * 5, time.Second * 6},
		},
		diffs: map[string]*diff{
			"metric1": {baseline: time.Second, count: 3, simSum: 6},
		},
		cached: map[string]int{
			"metric1": 1,
		},
	}

	writer := &bytes.Buffer{}
	store.WriteMetrics(writer)

	expected := `  cumulative  average  maximum  cache  percent  cached  total 
  cumulative  average  maximum  cache  percent  cached  total 
  cumulative  average  maximum  cache  percent  cached  total 
  duration  duration  duration  cache  cached  count  count  template
  ---------- -------- -------- --------- -------  ------  ----- --------
  3s 2s 3s 0 0.0 0 3 metric1
  15s 5s 6s 1 100.0 1 3 metric2
`

	if writer.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, writer.String())
	}
}
