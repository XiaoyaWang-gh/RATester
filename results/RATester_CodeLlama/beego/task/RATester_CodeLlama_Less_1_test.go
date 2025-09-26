package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &MapSorter{
		Keys: []string{"a", "b", "c"},
		Vals: []Tasker{
			&Task{
				Taskname: "a",
				Next:     time.Now().Add(time.Hour),
			},
			&Task{
				Taskname: "b",
				Next:     time.Now().Add(time.Hour * 2),
			},
			&Task{
				Taskname: "c",
				Next:     time.Now().Add(time.Hour * 3),
			},
		},
	}
	ms.Sort()
	if ms.Vals[0].GetNext(context.Background()).Before(ms.Vals[1].GetNext(context.Background())) {
		t.Errorf("ms.Vals[0].GetNext(context.Background()).Before(ms.Vals[1].GetNext(context.Background())) = %v, want %v", ms.Vals[0].GetNext(context.Background()).Before(ms.Vals[1].GetNext(context.Background())), false)
	}
	if ms.Vals[1].GetNext(context.Background()).Before(ms.Vals[2].GetNext(context.Background())) {
		t.Errorf("ms.Vals[1].GetNext(context.Background()).Before(ms.Vals[2].GetNext(context.Background())) = %v, want %v", ms.Vals[1].GetNext(context.Background()).Before(ms.Vals[2].GetNext(context.Background())), true)
	}
}
