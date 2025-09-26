package requestdecorator

import (
	"fmt"
	"testing"

	"github.com/miekg/dns"
)

func TestGetRecord_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	client := &dns.Client{}

	// CONTEXT_1
	msg := &dns.Msg{}

	// CONTEXT_2
	server := "8.8.8.8"
	port := "53"

	// CONTEXT_3
	msg.SetQuestion("example.com.", dns.TypeCNAME)

	// CONTEXT_4
	expected := &cnameResolv{
		TTL:    10,
		Record: "example.com.",
	}

	// CONTEXT_5
	server = "8.8.8.8"

	// CONTEXT_6
	port = "53"

	// CONTEXT_7
	actual, err := getRecord(client, msg, server, port)
	if err != nil {
		t.Fatalf("getRecord() error = %v", err)
	}

	// CONTEXT_8
	if actual.TTL != expected.TTL {
		t.Errorf("getRecord() TTL = %v, want %v", actual.TTL, expected.TTL)
	}

	// CONTEXT_9
	if actual.Record != expected.Record {
		t.Errorf("getRecord() Record = %v, want %v", actual.Record, expected.Record)
	}

	// CONTEXT_10
	if err != nil {
		t.Errorf("getRecord() err = %v, want nil", err)
	}

	// CONTEXT_11
	if actual != expected {
		t.Errorf("getRecord() = %v, want %v", actual, expected)
	}
}
