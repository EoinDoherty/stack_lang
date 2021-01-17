package main

import (
	"fmt"
	"stack_lang/parser"
	"stack_lang/runtime"
)

func main() {
	s := &runtime.Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println("hello world")

	tokens, err := parser.GetTokens("file.txt")

	for _, thing := range tokens {
		fmt.Printf("\"%s\"\n", thing)
	}
	fmt.Println(tokens)
	fmt.Println(err)

	fmt.Println("")
	fmt.Println("Pushing hello world")
	str := "hello world"
	runtime.PushString(str, s)
	runtime.PrintString(s)
}
