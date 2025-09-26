package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRun_2(t *testing.T) {
	type fields struct {
		Taskname string
		Spec     *Schedule
		SpecStr  string
		DoFunc   TaskFunc
		Prev     time.Time
		Next     time.Time
		Timeout  time.Duration
		Errlist  []*taskerr
		ErrLimit int
		errCnt   int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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

			task := &Task{
				Taskname: tt.fields.Taskname,
				Spec:     tt.fields.Spec,
				SpecStr:  tt.fields.SpecStr,
				DoFunc:   tt.fields.DoFunc,
				Prev:     tt.fields.Prev,
				Next:     tt.fields.Next,
				Timeout:  tt.fields.Timeout,
				Errlist:  tt.fields.Errlist,
				ErrLimit: tt.fields.ErrLimit,
				errCnt:   tt.fields.errCnt,
			}
			if err := task.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Task.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
