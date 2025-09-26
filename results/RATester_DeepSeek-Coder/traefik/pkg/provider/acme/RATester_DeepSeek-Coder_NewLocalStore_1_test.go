package acme

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLocalStore_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want *LocalStore
	}{
		{
			name: "TestNewLocalStore",
			args: args{filename: "test.db"},
			want: &LocalStore{
				filename:     "test.db",
				saveDataChan: make(chan map[string]*StoredData),
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

			got := NewLocalStore(tt.args.filename)
			if got.filename != tt.want.filename {
				t.Errorf("NewLocalStore() = %v, want %v", got.filename, tt.want.filename)
			}
			if reflect.TypeOf(got.saveDataChan) != reflect.TypeOf(tt.want.saveDataChan) {
				t.Errorf("NewLocalStore() = %v, want %v", reflect.TypeOf(got.saveDataChan), reflect.TypeOf(tt.want.saveDataChan))
			}
		})
	}
}
