package task

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetAllTasks_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []Tasker
	}{
		{
			name: "TestGetAllTasks",
			args: args{
				ctx: context.Background(),
			},
			want: []Tasker{
				&Task{
					Taskname: "Task1",
					Spec:     &Schedule{},
					SpecStr:  "* * * * *",
					Prev:     time.Now(),
					Next:     time.Now().Add(time.Minute),
					Timeout:  time.Minute,
				},
				&Task{
					Taskname: "Task2",
					Spec:     &Schedule{},
					SpecStr:  "* * * * *",
					Prev:     time.Now(),
					Next:     time.Now().Add(time.Minute),
					Timeout:  time.Minute,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := GetAllTasks()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
