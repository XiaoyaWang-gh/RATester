package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOrmStub_1(t *testing.T) {
	tests := []struct {
		name string
		want *OrmStub
	}{
		{
			name: "Test newOrmStub",
			want: &OrmStub{
				ms: make([]*Mock, 0, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newOrmStub(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOrmStub() = %v, want %v", got, tt.want)
			}
		})
	}
}
