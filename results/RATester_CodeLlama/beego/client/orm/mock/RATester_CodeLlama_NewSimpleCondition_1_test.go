package mock

import (
	"fmt"
	"testing"
)

func TestNewSimpleCondition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tableName := "tableName"
	methodName := "methodName"
	condition := NewSimpleCondition(tableName, methodName)
	if condition.(*SimpleCondition).tableName != tableName {
		t.Errorf("NewSimpleCondition() tableName = %v, want %v", condition.(*SimpleCondition).tableName, tableName)
	}
	if condition.(*SimpleCondition).method != methodName {
		t.Errorf("NewSimpleCondition() method = %v, want %v", condition.(*SimpleCondition).method, methodName)
	}
}
