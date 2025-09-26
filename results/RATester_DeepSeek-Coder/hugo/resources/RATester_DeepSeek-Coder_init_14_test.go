package resources

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
)

func TestInit_14(t *testing.T) {
	type fields struct {
		value    uint64
		size     int64
		initOnce sync.Once
	}
	type args struct {
		l hugio.ReadSeekCloserProvider
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &resourceHash{
				value:    tt.fields.value,
				size:     tt.fields.size,
				initOnce: tt.fields.initOnce,
			}
			if err := r.init(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("resourceHash.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
