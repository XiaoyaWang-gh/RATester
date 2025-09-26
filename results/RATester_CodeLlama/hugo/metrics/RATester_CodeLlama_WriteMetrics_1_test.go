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

	s := &Store{
		calculateHints: true,
		metrics: map[string][]time.Duration{
			"foo": {1, 2, 3, 4, 5},
			"bar": {1, 2, 3, 4, 5},
			"baz": {1, 2, 3, 4, 5},
		},
		diffs: map[string]*diff{
			"foo": {baseline: 1, count: 5},
			"bar": {baseline: 1, count: 5},
			"baz": {baseline: 1, count: 5},
		},
		cached: map[string]int{
			"foo": 1,
			"bar": 1,
			"baz": 1,
		},
	}

	var buf bytes.Buffer
	s.WriteMetrics(&buf)

	expected := `  cumulative  average  maximum  cache  percent  cached  total  template
  duration  duration  duration  potential  cached  count  count  template
  ---------  --------  --------  ---------  ------  ------  ------  ---------
  15        3         5         0          0       5      5       foo
  15        3         5         0          0       5      5       bar
  15        3         5         0          0       5      5       baz
`

	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
