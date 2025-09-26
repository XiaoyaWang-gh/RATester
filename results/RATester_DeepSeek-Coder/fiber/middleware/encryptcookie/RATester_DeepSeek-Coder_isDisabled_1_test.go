package encryptcookie

import (
	"fmt"
	"testing"
)

func TestIsDisabled_1(t *testing.T) {
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
				key:    "test1",
				except: []string{"test1", "test2"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				key:    "test3",
				except: []string{"test1", "test2"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				key:    "test1",
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
