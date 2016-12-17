package test

import (
	"testing"
)

//--- FAIL: TestA (0.00s)
//	test_test.go:13: got B want A
//--- FAIL: TestOne (0.00s)
//	test_test.go:23: got 2 want 1
//	test_test.go:29: got 2 want 1
//FAIL

func TestA(t *testing.T) {
	actual := A()
	expected := "A"
	if actual != expected {
		//--- FAIL: TestA (0.00s)
		//	test_test.go:11: got B want A
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestOne(t *testing.T) {
	actual := One()
	expected := 1
	if actual != expected {
		//--- FAIL: TestOne (0.00s)
		//	test_test.go:23: got 2 want 1
		t.Errorf("got %v want %v", actual, expected)
	}

	if actual != expected {
		//	test_test.go:23: got 2 want 1
		t.Errorf("got %v want %v", actual, expected)
	}
}
