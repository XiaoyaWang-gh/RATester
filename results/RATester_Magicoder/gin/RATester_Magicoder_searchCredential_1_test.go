package gin

import (
	"fmt"
	"testing"
)

func TestsearchCredential_1(t *testing.T) {
	type args struct {
		authValue string
	}
	tests := []struct {
		name  string
		a     authPairs
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Test Case 1",
			a: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			args:  args{authValue: "value1"},
			want:  "user1",
			want1: true,
		},
		{
			name: "Test Case 2",
			a: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			args:  args{authValue: "value3"},
			want:  "",
			want1: false,
		},
		{
			name: "Test Case 3",
			a: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			args:  args{authValue: ""},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := tt.a.searchCredential(tt.args.authValue)
			if got != tt.want {
				t.Errorf("searchCredential() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("searchCredential() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
