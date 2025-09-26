package auth

import (
	"fmt"
	"testing"
)

func TestLoadUsers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var fileName string
	var appendUsers []string
	var users []string
	var err error

	// Test case 1:
	fileName = "testdata/users.txt"
	appendUsers = []string{"user1", "user2"}
	users, err = loadUsers(fileName, appendUsers)
	if err != nil {
		t.Errorf("loadUsers() error = %v", err)
		return
	}
	if len(users) != 3 {
		t.Errorf("loadUsers() users = %v, want %v", users, appendUsers)
	}

	// Test case 2:
	fileName = ""
	appendUsers = []string{"user1", "user2"}
	users, err = loadUsers(fileName, appendUsers)
	if err != nil {
		t.Errorf("loadUsers() error = %v", err)
		return
	}
	if len(users) != 2 {
		t.Errorf("loadUsers() users = %v, want %v", users, appendUsers)
	}

	// Test case 3:
	fileName = "testdata/users.txt"
	appendUsers = []string{}
	users, err = loadUsers(fileName, appendUsers)
	if err != nil {
		t.Errorf("loadUsers() error = %v", err)
		return
	}
	if len(users) != 2 {
		t.Errorf("loadUsers() users = %v, want %v", users, appendUsers)
	}

	// Test case 4:
	fileName = ""
	appendUsers = []string{}
	users, err = loadUsers(fileName, appendUsers)
	if err != nil {
		t.Errorf("loadUsers() error = %v", err)
		return
	}
	if len(users) != 0 {
		t.Errorf("loadUsers() users = %v, want %v", users, appendUsers)
	}
}
