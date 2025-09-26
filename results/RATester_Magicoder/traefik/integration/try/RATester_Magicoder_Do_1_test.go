package try

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDo_1(t *testing.T) {
	type args struct {
		timeout   time.Duration
		operation DoCondition
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_case_1",
			args: args{
				timeout:   1 * time.Second,
				operation: func() error { return nil },
			},
			wantErr: false,
		},
		{
			name: "test_case_2",
			args: args{
				timeout:   1 * time.Second,
				operation: func() error { return errors.New("test error") },
			},
			wantErr: true,
		},
		{
			name: "test_case_3",
			args: args{
				timeout:   0,
				operation: func() error { return nil },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := Do(tt.args.timeout, tt.args.operation); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
