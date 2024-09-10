package test_test

import (
	"demo_test/tool"
	"testing"
)

// testing.T是testing包提供的用于方便测试的结构体
func TestSum(t *testing.T) {
	a, b := 10, 101
	expected := 111

	actual := tool.SumInt(a, b)
	if actual != expected {
		t.Errorf("Sum(%d,%d) expected %d,actual is %d", a, b, expected, actual)
	}
}

func TestEqual(t *testing.T) {
	a, b := 10, 101
	expected := false

	actual := tool.Equal(a, b)
	if actual != expected {
		t.Errorf("Sum(%d,%d) expected %t,actual is %t", a, b, expected, actual)
	}
}
