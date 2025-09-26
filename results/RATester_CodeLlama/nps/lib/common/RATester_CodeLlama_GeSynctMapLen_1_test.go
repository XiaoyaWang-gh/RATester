package common

import (
	"fmt"
	"sync"
	"testing"
)

func TestGeSynctMapLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := sync.Map{}
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)
	m.Store("d", 4)
	m.Store("e", 5)
	m.Store("f", 6)
	m.Store("g", 7)
	m.Store("h", 8)
	m.Store("i", 9)
	m.Store("j", 10)
	m.Store("k", 11)
	m.Store("l", 12)
	m.Store("m", 13)
	m.Store("n", 14)
	m.Store("o", 15)
	m.Store("p", 16)
	m.Store("q", 17)
	m.Store("r", 18)
	m.Store("s", 19)
	m.Store("t", 20)
	m.Store("u", 21)
	m.Store("v", 22)
	m.Store("w", 23)
	m.Store("x", 24)
	m.Store("y", 25)
	m.Store("z", 26)
	if got := GeSynctMapLen(m); got != 26 {
		t.Errorf("GeSynctMapLen() = %v, want %v", got, 26)
	}
}
