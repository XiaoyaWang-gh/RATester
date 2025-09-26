package compare

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew_66(t *testing.T) {
	tests := []struct {
		name string
		loc  *time.Location
		want *Namespace
	}{
		{
			name: "case 1",
			loc:  time.UTC,
			want: &Namespace{
				loc: time.UTC,
			},
		},
		{
			name: "case 2",
			loc:  time.Local,
			want: &Namespace{
				loc: time.Local,
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

			if got := New(tt.loc, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
