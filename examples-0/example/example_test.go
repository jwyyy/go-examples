package example_test

import (
	"testing"

	"example.com/go-demo-0/example"
)

func TestFunction(t *testing.T) {
	if example.Test() != "Go testing..." {
		t.Fatal("Wrong example.Test()")
	}
}
