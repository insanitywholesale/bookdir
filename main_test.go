package main

import (
	"testing"
)

func TestGetServiceName(t *testing.T) {
	expectedResult := "bookdir"
	actualResult := getServiceName()
	if expectedResult != actualResult {
		t.Fatal("service name is wrong")
	}
	t.Log("service name is:", actualResult)
}
