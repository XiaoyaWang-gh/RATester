package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		SavePath: "localhost:6379",
		Poolsize: 10,
		Password: "password",
		DbNum:    0,
	}

	err := rp.SessionInit(ctx, 100, `{"save_path": "localhost:6379", "poolsize": 10, "password": "password", "db_num": 0}`)
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if rp.SavePath != "localhost:6379" {
		t.Errorf("Expected SavePath to be 'localhost:6379', but got %s", rp.SavePath)
	}

	if rp.Poolsize != 10 {
		t.Errorf("Expected Poolsize to be 10, but got %d", rp.Poolsize)
	}

	if rp.Password != "password" {
		t.Errorf("Expected Password to be 'password', but got %s", rp.Password)
	}

	if rp.DbNum != 0 {
		t.Errorf("Expected DbNum to be 0, but got %d", rp.DbNum)
	}

	if rp.poollist == nil {
		t.Error("Expected poollist to be initialized, but it is nil")
	}
}
