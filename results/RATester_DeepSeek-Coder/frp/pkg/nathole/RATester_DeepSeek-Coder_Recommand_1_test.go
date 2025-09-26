package nathole

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRecommand_1(t *testing.T) {
	t.Parallel()

	type fields struct {
		mu             sync.Mutex
		scores         []*BehaviorScore
		LastUpdateTime time.Time
	}

	tests := []struct {
		name   string
		fields fields
		want   int
		want1  int
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mhr := &MakeHoleRecords{
				mu:             tt.fields.mu,
				scores:         tt.fields.scores,
				LastUpdateTime: tt.fields.LastUpdateTime,
			}
			got, got1 := mhr.Recommand()
			if got != tt.want {
				t.Errorf("Recommand() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Recommand() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
