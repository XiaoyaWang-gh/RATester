package udp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWRRLoadBalancer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	want := &WRRLoadBalancer{
		index: -1,
	}
	got := NewWRRLoadBalancer()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("NewWRRLoadBalancer() = %v, want %v", got, want)
	}
}
