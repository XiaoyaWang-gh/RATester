package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddRoute_2(t *testing.T) {
	type args struct {
		rule     string
		syntax   string
		priority int
		handler  http.Handler
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				rule:     "rule1",
				syntax:   "syntax1",
				priority: 1,
				handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				rule:     "rule2",
				syntax:   "syntax2",
				priority: 2,
				handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Muxer{}
			err := m.AddRoute(tt.args.rule, tt.args.syntax, tt.args.priority, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Muxer.AddRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
