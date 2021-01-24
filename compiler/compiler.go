package compiler

import (
	"fmt"
	"io"
	"os"
	"stack_lang/parser"
	"stack_lang/runtime"
	"strconv"
)

func CompileFile(inputFilename string, outputFilename string) error {
    tokens, err := parser.GetTokens(inputFilename)

    if err != nil && err != io.EOF {
        return fmt.Errorf("CompileFile input error: %v", err)
    }

    out, err := os.Create(outputFilename)

    if err != nil {
        return fmt.Errorf("CompileFile output file error: %v", err)
    }
    defer out.Close()

    return compile(tokens, out)
}

func compile(tokens []string, out io.Writer) error {
    var err error = nil

    for {
        tokens, err = compileTokens(tokens, out)

        if err != nil {
            return fmt.Errorf("CompileFile: %v", err)
        }

        if len(tokens) == 0 {
            break;
        }
    }

    return nil
}

func compileTokens(tokens []string, out io.Writer) ([]string, error) {
    if len(tokens) == 0 {
        return tokens, nil
    }

    first := tokens[0]

    switch first {
        case "push":
            return compilePush(tokens[1:], out)
        case "pop":
            return tokens[1:], pushOpCode(runtime.POP, out)
        case "pops":
            return tokens[1:], pushOpCode(runtime.POP_STR, out)
        case "add":
            return tokens[1:], pushOpCode(runtime.ADD, out)
        case "sub":
            return tokens[1:], pushOpCode(runtime.SUB, out)
        case "div":
            return tokens[1:], pushOpCode(runtime.DIV, out)
        case "mul":
            return tokens[1:], pushOpCode(runtime.MUL, out)
        case "print":
            return tokens[1:], pushOpCode(runtime.PRINT_VAL, out)
        case "prints":
            return tokens[1:], pushOpCode(runtime.PRINT_STR, out)
        default:
            return tokens, fmt.Errorf("compileTokens: Unrecognized token: " + first)
    }
}

func compilePush(tokens []string, out io.Writer) ([]string, error) {
    if len(tokens) == 0 {
        return tokens, fmt.Errorf("Argument not specified for push")
    }

    if isString(tokens[0]) {
        return compilePushStr(tokens, out)
    }

    int8Rep, err := strconv.ParseUint(tokens[0], 10, 8)

    if err != nil {
        return tokens[1:], err
    }

    _, err = out.Write([]byte{runtime.PUSH, byte(int8Rep)})

    return tokens[1:], err
}

func compilePushStr(tokens []string, out io.Writer) ([]string, error) {
    strBytes := []byte(unquoteString(tokens[0]))
    strBytes = append(strBytes, runtime.END_STR)

    for i := len(strBytes) - 1; i >= 0; i-- {
    // for _,b := range strBytes {
        _, err := out.Write([]byte{runtime.PUSH, strBytes[i]})

        if err != nil {
            return tokens, err
        }
    }

    return tokens[1:], nil
}

func pushOpCode(code byte, out io.Writer) error {
    _, err := out.Write([]byte{code})
    return err
}

func isString(token string) bool {
    length := len(token)

    if length < 2 {
        return false
    }

    return token[0] == '"' && token[length-1] == '"'
}

func unquoteString(quotedString string) string {
    return quotedString[1:len(quotedString)-1]
}
