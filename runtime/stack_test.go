package runtime

import "testing"

func TestPush(t *testing.T) {
	s := &Stack{}
	checkData(t, s, []byte{})

	s.Push(1)
	s.Push(2)
	checkData(t, s, []byte{1, 2})
}

func TestPop(t *testing.T) {
	expected := []byte{1, 2, 3}
	s := &Stack{data: expected}

	checkData(t, s, []byte{1, 2, 3})

	for i := len(expected) - 1; i >= 0; i-- {
		val := s.Pop()

		if val != expected[i] {
			t.Errorf("Error: %d != %d", val, expected[i])
		}
	}
	checkData(t, s, []byte{})

	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop did not panic for empty stack")
		}
	}()

	s = &Stack{}
	s.Pop()
}

func TestIsEmpty(t *testing.T) {
	s := &Stack{}

	if !s.IsEmpty() {
		t.Error("IsEmpty=false for empty stack")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Error("IsEmpty=true for non-empty stack")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Error("IsEmpty=false for empty stack")
	}
}

func checkData(t *testing.T, s *Stack, expected []byte) {
	ldata := len(s.data)
	lexpected := len(expected)
	if lexpected != ldata {
		t.Errorf("Data length mismatch: %d != %d", lexpected, ldata)
	}

	for i, elem := range s.data {
		expectedI := expected[i]
		if elem != expectedI {
			t.Errorf("Error: %d != %d", elem, expectedI)
		}
	}
}
