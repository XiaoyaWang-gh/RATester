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
			name: "success",
			args: args{
				timeout: 1 * time.Second,
				operation: func() error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				timeout: 1 * time.Second,
				operation: func() error {
					return errors.New("operation failed")
				},
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

			err := Do(tt.args.timeout, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
