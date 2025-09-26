package orm

import (
	"fmt"
	"reflect"
	"testing"

	lru "github.com/hashicorp/golang-lru"
)

func TestnewStmtDecoratorLruWithEvict_1(t *testing.T) {
	type args struct {
		cacheSize int
	}
	tests := []struct {
		name    string
		args    args
		want    *lru.Cache
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				cacheSize: 10,
			},
			want:    &lru.Cache{},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				cacheSize: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				cacheSize: -1,
			},
			want:    nil,
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

			got, err := newStmtDecoratorLruWithEvict(tt.args.cacheSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStmtDecoratorLruWithEvict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStmtDecoratorLruWithEvict() = %v, want %v", got, tt.want)
			}
		})
	}
}
