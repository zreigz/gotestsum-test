package main

import (
	"fmt"
	"testing"
)

func Test_Digging(t *testing.T) {
	fmt.Print("Hello")
}

func Test_Failing(t *testing.T) {
	t.Fatal("This test fails")
}
