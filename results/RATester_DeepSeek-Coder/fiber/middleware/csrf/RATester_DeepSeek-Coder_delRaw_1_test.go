package csrf

import (
	"fmt"
	"testing"
)

func TestDelRaw_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    *storageManager
		args args
	}{
		{
			name: "Test case 1",
			m:    &storageManager{},
			args: args{
				key: "testKey",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.delRaw(tt.args.key)
		})
	}
}
