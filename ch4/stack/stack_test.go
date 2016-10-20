// Write a simple unit test for the stack package. You should at least test
// that a Pop works after a Push.
package stack

import "testing"

func TestPushThenPop(t *testing.T) {
	s := Stack{}
	err := s.Push(1)
	if err != nil {
		t.Log("failed to push to stack.")
		t.FailNow()
	}
	v, err := s.Pop()
	if err != nil {
		t.Log("failed to pop from stack.")
		t.FailNow()
	}
	if v != 1 {
		t.Log("popped value != pushed value")
		t.Fail()
	}
}
