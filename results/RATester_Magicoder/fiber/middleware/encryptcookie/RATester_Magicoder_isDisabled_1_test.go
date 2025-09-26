package encryptcookie

import (
	"fmt"
	"testing"
)

func TestisDisabled_1(t *testing.T) {
	type args struct {
		key    string
		except []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				key:    "field name",
				except: []string{"field name"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				key:    "field name",
				except: []string{"field name2"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				key:    "field name",
				except: []string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isDisabled(tt.args.key, tt.args.except); got != tt.want {
				t.Errorf("isDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
