package tcp

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestNext_2(t *testing.T) {
	type fields struct {
		servers       []server
		lock          sync.Mutex
		currentWeight int
		index         int
	}
	tests := []struct {
		name    string
		fields  fields
		want    Handler
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

			b := &WRRLoadBalancer{
				servers:       tt.fields.servers,
				lock:          tt.fields.lock,
				currentWeight: tt.fields.currentWeight,
				index:         tt.fields.index,
			}
			got, err := b.next()
			if (err != nil) != tt.wantErr {
				t.Errorf("WRRLoadBalancer.next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WRRLoadBalancer.next() = %v, want %v", got, tt.want)
			}
		})
	}
}
