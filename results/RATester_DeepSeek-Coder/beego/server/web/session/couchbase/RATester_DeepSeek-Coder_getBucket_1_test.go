package couchbase

import (
	"fmt"
	"reflect"
	"testing"

	couchbase "github.com/couchbase/go-couchbase"
)

func TestGetBucket_1(t *testing.T) {
	type fields struct {
		maxlifetime int64
		SavePath    string
		Pool        string
		Bucket      string
		b           *couchbase.Bucket
	}
	tests := []struct {
		name   string
		fields fields
		want   *couchbase.Bucket
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

			cp := &Provider{
				maxlifetime: tt.fields.maxlifetime,
				SavePath:    tt.fields.SavePath,
				Pool:        tt.fields.Pool,
				Bucket:      tt.fields.Bucket,
				b:           tt.fields.b,
			}
			if got := cp.getBucket(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.getBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
