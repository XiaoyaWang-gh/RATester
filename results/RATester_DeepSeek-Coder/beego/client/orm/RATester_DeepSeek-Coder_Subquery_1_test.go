package orm

import (
	"fmt"
	"testing"
)

func TestSubquery_1(t *testing.T) {
	type args struct {
		sub   string
		alias string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				sub:   "SELECT * FROM table",
				alias: "t",
			},
			want: "(SELECT * FROM table) AS t",
		},
		{
			name: "Test case 2",
			args: args{
				sub:   "SELECT * FROM table2",
				alias: "t2",
			},
			want: "(SELECT * FROM table2) AS t2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			qb := &MySQLQueryBuilder{}
			if got := qb.Subquery(tt.args.sub, tt.args.alias); got != tt.want {
				t.Errorf("Subquery() = %v, want %v", got, tt.want)
			}
		})
	}
}
