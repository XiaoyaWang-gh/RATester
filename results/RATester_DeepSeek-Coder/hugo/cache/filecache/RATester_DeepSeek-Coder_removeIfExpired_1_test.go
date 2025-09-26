package filecache

import (
	"fmt"
	"testing"
)

func TestRemoveIfExpired_1(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       *Cache
		args    args
		want    bool
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

			got, err := tt.c.removeIfExpired(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeIfExpired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("removeIfExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
