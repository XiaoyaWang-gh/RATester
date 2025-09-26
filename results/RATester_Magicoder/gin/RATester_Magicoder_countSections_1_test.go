package gin

import (
	"fmt"
	"testing"
)

func TestcountSections_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Test case 1",
			args: args{path: "/home/user/documents/file.txt"},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{path: "/"},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{path: "/usr/local/bin"},
			want: 3,
		},
		{
			name: "Test case 4",
			args: args{path: "/etc/apache2/sites-available/"},
			want: 4,
		},
		{
			name: "Test case 5",
			args: args{path: "/var/log/syslog"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := countSections(tt.args.path); got != tt.want {
				t.Errorf("countSections() = %v, want %v", got, tt.want)
			}
		})
	}
}
