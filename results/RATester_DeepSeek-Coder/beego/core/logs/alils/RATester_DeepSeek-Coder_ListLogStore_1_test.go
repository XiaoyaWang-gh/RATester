package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListLogStore_1(t *testing.T) {
	type fields struct {
		Name            string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
	}
	tests := []struct {
		name       string
		fields     fields
		wantStores []string
		wantErr    bool
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

			p := &LogProject{
				Name:            tt.fields.Name,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
			}
			gotStores, err := p.ListLogStore()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListLogStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStores, tt.wantStores) {
				t.Errorf("ListLogStore() gotStores = %v, want %v", gotStores, tt.wantStores)
			}
		})
	}
}
