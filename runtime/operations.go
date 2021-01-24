package runtime

import "fmt"

func Add(s *Stack) {
	x := s.Pop()
	y := s.Pop()

	s.Push(x + y)
}

func Subtract(s *Stack) {
	x := s.Pop()
	y := s.Pop()

	s.Push(x - y)
}

func Multiply(s *Stack) {
	x := s.Pop()
	y := s.Pop()

	s.Push(x * y)
}

func Divide(s *Stack) {
	x := s.Pop()
	y := s.Pop()

	s.Push(x / y)
}

func PrintVal(s *Stack) {
    val := s.Pop()
	fmt.Println(val)
    s.Push(val)
}

func PrintString(s *Stack) {
	str := FetchStrBytes(s)
	fmt.Printf("%s\n", str)
    PushString(str, s)
}

func FetchStrBytes(s *Stack) []byte {
	str_bytes := make([]byte, 1)
    counter := 0

	for {
        counter++
		v := s.Pop()
		if v == END_STR {
			break
		}
		str_bytes = append(str_bytes, v)
	}

	return str_bytes
}

func PushString(str []byte, s *Stack) {
	s.Push(END_STR)

	for i := len(str) - 1; i >= 0; i-- {
		s.Push(str[i])
	}
}
