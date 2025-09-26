package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		maxlifetime:           10,
		SavePath:              "localhost:26379;localhost:26380;localhost:26381",
		Poolsize:              10,
		Password:              "password",
		DbNum:                 0,
		IdleTimeoutStr:        "10s",
		IdleCheckFrequencyStr: "10s",
		MaxRetries:            3,
	}

	err := rp.SessionInit(ctx, 10, `{"save_path":"localhost:26379;localhost:26380;localhost:26381","poolsize":10,"password":"password","db_num":0,"idle_timeout":"10s","idle_check_frequency":"10s","max_retries":3}`)
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}

	if rp.poollist == nil {
		t.Errorf("SessionInit failed: poollist is nil")
	}

	if rp.poollist.Ping(ctx).Err() != nil {
		t.Errorf("SessionInit failed: Ping failed")
	}
}
