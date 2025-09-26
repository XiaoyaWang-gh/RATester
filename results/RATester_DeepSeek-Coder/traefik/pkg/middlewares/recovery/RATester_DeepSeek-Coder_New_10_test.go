package recovery

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNew_10(t *testing.T) {
	type args struct {
		ctx  context.Context
		next http.Handler
	}

	tests := []struct {
		name    string
		args    args
		want    http.Handler
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:  context.Background(),
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			want:    &recovery{next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
			wantErr: false,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := New(tt.args.ctx, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
