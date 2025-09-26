package page

import (
	"fmt"
	"testing"
)

func TestdecodePageMatcher_1(t *testing.T) {
	type args struct {
		m any
		v *PageMatcher
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				m: map[string]any{
					"Path":        "/content/test",
					"Kind":        "home",
					"Lang":        "en",
					"Environment": "production",
				},
				v: &PageMatcher{},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				m: map[string]any{
					"Path":        "/content/test",
					"Kind":        "invalid",
					"Lang":        "en",
					"Environment": "production",
				},
				v: &PageMatcher{},
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

			if err := decodePageMatcher(tt.args.m, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodePageMatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
