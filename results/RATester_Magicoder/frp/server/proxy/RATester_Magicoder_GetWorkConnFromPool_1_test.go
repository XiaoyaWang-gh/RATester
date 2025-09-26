package proxy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
)

func TestGetWorkConnFromPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pxy := &BaseProxy{
		ctx: ctx,
		getWorkConnFn: func() (net.Conn, error) {
			return nil, errors.New("failed to get work connection")
		},
	}

	src, dst := net.IPAddr{}, net.IPAddr{}
	workConn, err := pxy.GetWorkConnFromPool(&src, &dst)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if workConn != nil {
		t.Errorf("Expected nil workConn, but got %v", workConn)
	}
}
