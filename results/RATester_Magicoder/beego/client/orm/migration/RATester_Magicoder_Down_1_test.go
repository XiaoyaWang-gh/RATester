package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDown_1(t *testing.T) {
	tests := []struct {
		name     string
		m        *Migration
		want     *Migration
		wantSqls []string
		wantErr  bool
	}{
		{
			name: "Test alter",
			m: &Migration{
				ModifyType: "alter",
			},
			want: &Migration{
				ModifyType: "reverse",
			},
			wantSqls: []string{},
			wantErr:  false,
		},
		{
			name: "Test create",
			m: &Migration{
				ModifyType: "create",
			},
			want: &Migration{
				ModifyType: "delete",
			},
			wantSqls: []string{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.Down()
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Down() = %v, want %v", tt.m, tt.want)
			}
			if !reflect.DeepEqual(tt.m.sqls, tt.wantSqls) {
				t.Errorf("Down() = %v, want %v", tt.m.sqls, tt.wantSqls)
			}
		})
	}
}
