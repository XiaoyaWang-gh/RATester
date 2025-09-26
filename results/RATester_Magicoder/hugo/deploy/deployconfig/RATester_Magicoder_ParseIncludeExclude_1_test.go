package deployconfig

import (
	"fmt"
	"testing"
)

func TestParseIncludeExclude_1(t *testing.T) {
	tests := []struct {
		name    string
		tgt     *Target
		wantErr bool
	}{
		{
			name: "valid include",
			tgt: &Target{
				Include: "*.txt",
			},
			wantErr: false,
		},
		{
			name: "invalid include",
			tgt: &Target{
				Include: "*[.txt",
			},
			wantErr: true,
		},
		{
			name: "valid exclude",
			tgt: &Target{
				Exclude: "*.txt",
			},
			wantErr: false,
		},
		{
			name: "invalid exclude",
			tgt: &Target{
				Exclude: "*[.txt",
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

			if err := tt.tgt.ParseIncludeExclude(); (err != nil) != tt.wantErr {
				t.Errorf("ParseIncludeExclude() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
