package redis

import (
	"fmt"
	"testing"
)

func TestStartAndGC_5(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		rc      *Cache
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

			rc := &Cache{}
			if err := rc.StartAndGC(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Cache.StartAndGC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
