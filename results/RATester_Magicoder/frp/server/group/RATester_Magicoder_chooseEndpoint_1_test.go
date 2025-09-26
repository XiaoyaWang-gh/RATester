package group

import (
	"fmt"
	"testing"
)

func TestChooseEndpoint_1(t *testing.T) {
	g := &HTTPGroup{
		group:           "testGroup",
		groupKey:        "testGroupKey",
		domain:          "testDomain",
		location:        "testLocation",
		routeByHTTPUser: "testRouteByHTTPUser",
		pxyNames:        []string{"proxy1", "proxy2", "proxy3"},
	}

	tests := []struct {
		name    string
		g       *HTTPGroup
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			g:       g,
			want:    "proxy1",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			g:       g,
			want:    "proxy2",
			wantErr: false,
		},
		{
			name:    "Test case 3",
			g:       g,
			want:    "proxy3",
			wantErr: false,
		},
		{
			name:    "Test case 4",
			g:       g,
			want:    "",
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

			got, err := tt.g.chooseEndpoint()
			if (err != nil) != tt.wantErr {
				t.Errorf("chooseEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("chooseEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
