package adapter

import "testing"

var expect = "adapteeimpl"

func TestAdapter(t *testing.T) {
	adaptee := NewAdatpee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
