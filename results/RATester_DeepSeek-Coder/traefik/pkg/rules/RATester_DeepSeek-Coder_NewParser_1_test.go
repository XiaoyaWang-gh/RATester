package rules

import (
	"fmt"
	"testing"
)

func TestNewParser_1(t *testing.T) {
	type args struct {
		matchers []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewParser_Success",
			args: args{
				matchers: []string{"matcher1", "matcher2"},
			},
			wantErr: false,
		},
		{
			name: "TestNewParser_Failure",
			args: args{
				matchers: []string{"matcher3", "matcher4"},
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

			_, err := NewParser(tt.args.matchers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
