package httpcache

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCompile_4(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		want    ConfigCompiled
		wantErr bool
	}{
		{
			name: "Test case 1",
			config: Config{
				Cache: Cache{
					For: GlobMatcher{
						Includes: []string{"*"},
					},
				},
				Polls: []PollConfig{
					{
						For: GlobMatcher{
							Includes: []string{"*"},
						},
						Disable: false,
						Low:     1 * time.Minute,
						High:    10 * time.Minute,
					},
				},
			},
			want: ConfigCompiled{
				For: func(s string) bool {
					return true
				},
				PollConfigs: []PollConfigCompiled{
					{
						For: func(s string) bool {
							return true
						},
						Config: PollConfig{
							For: GlobMatcher{
								Includes: []string{"*"},
							},
							Disable: false,
							Low:     1 * time.Minute,
							High:    10 * time.Minute,
						},
					},
				},
			},
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

			got, err := tt.config.Compile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}
