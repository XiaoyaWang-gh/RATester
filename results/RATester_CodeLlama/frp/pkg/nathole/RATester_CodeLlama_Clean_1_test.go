package nathole

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestClean_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Analyzer{
		records:             make(map[string]*MakeHoleRecords),
		dataReserveDuration: time.Second * 10,
	}

	a.records["1.1.1.1:10000:2.2.2.2:20000"] = &MakeHoleRecords{
		scores: []*BehaviorScore{
			{Mode: 1, Index: 1, Score: 10},
			{Mode: 1, Index: 2, Score: 10},
			{Mode: 1, Index: 3, Score: 10},
			{Mode: 1, Index: 4, Score: 10},
			{Mode: 1, Index: 5, Score: 10},
			{Mode: 1, Index: 6, Score: 10},
			{Mode: 1, Index: 7, Score: 10},
			{Mode: 1, Index: 8, Score: 10},
			{Mode: 1, Index: 9, Score: 10},
			{Mode: 1, Index: 10, Score: 10},
		},
	}

	a.records["1.1.1.1:10000:2.2.2.2:20000"].LastUpdateTime = time.Now().Add(-time.Second * 11)

	count, total := a.Clean()
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, total)

	_, ok := a.records["1.1.1.1:10000:2.2.2.2:20000"]
	assert.False(t, ok)
}
