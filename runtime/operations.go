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
	fmt.Println(s.Pop())
}

func PrintString(s *Stack) {
	str := FetchStrBytes(s)
	fmt.Printf("%s\n", str)
}

func FetchStrBytes(s *Stack) []byte {
	str_bytes := make([]byte, 1)

	for {
		v := s.Pop()
		if v == END_STR {
			break
		}
		str_bytes = append(str_bytes, v)
	}

	return str_bytes
}

func PushString(str string, s *Stack) {
	bs := []byte(str)

	s.Push(END_STR)

	for i := len(bs) - 1; i >= 0; i-- {
		s.Push(bs[i])
	}
}
