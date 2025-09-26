package controllers

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestclearIprecord_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ipRecord = sync.Map{}
	ipRecord.Store("192.168.1.1", &record{hasLoginFailTimes: 0, lastLoginTime: time.Now().Add(-time.Minute)})
	ipRecord.Store("192.168.1.2", &record{hasLoginFailTimes: 0, lastLoginTime: time.Now()})
	clearIprecord()
	_, ok := ipRecord.Load("192.168.1.1")
	if ok {
		t.Error("192.168.1.1 should be deleted")
	}
	_, ok = ipRecord.Load("192.168.1.2")
	if !ok {
		t.Error("192.168.1.2 should not be deleted")
	}
}
