package runtime

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var optable = map[byte]func(*Stack, *bufio.Reader) error {
    PUSH:      runPush,
    POP:       runPop,
    POP_STR:   runPopStr,
    ADD:       runAdd,
    SUB:       runSub,
    DIV:       runDiv,
    MUL:       runMul,
    PRINT_VAL: runPrintVal,
    PRINT_STR: runPrintStr,
}

func RunFile(filename string) error {
    file, err := os.Open(filename)

    if err != nil {
        return err
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    err = interpretBytes(reader)

    if err == io.EOF {
        return nil
    }

    return err
}

func interpretBytes(reader *bufio.Reader) error {
    s := &Stack{}

    for {
        b, err := reader.ReadByte()

        if err != nil {
            return err
        }

        f, ok := optable[b]

        if !ok {
            return fmt.Errorf("Unknown opcode: %v", b)
        }

        err = f(s, reader)
        if err != nil {
            return err
        }
    }
}

func runPush(s *Stack, reader *bufio.Reader) error {
    data, err := reader.ReadByte()

    if err != nil {
        return err
    }

    s.Push(data)
    return nil
}

func runPop(s *Stack, reader *bufio.Reader) error {
    s.Pop()
    return nil
}

func runPopStr(s *Stack, reader *bufio.Reader) error {
    FetchStrBytes(s)
    return nil
}

func runAdd(s *Stack, reader *bufio.Reader) error {
    Add(s)
    return nil
}

func runSub(s *Stack, reader *bufio.Reader) error {
    Subtract(s)
    return nil
}

func runDiv(s *Stack, reader *bufio.Reader) error {
    Divide(s)
    return nil
}

func runMul(s *Stack, reader *bufio.Reader) error {
    Multiply(s)
    return nil
}

func runPrintVal(s *Stack, reader *bufio.Reader) error {
    PrintVal(s)
    return nil
}

func runPrintStr(s *Stack, reader *bufio.Reader) error {
    PrintString(s)
    return nil
}

