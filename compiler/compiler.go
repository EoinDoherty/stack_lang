package compiler

import (
	"fmt"
	"io"
	"stack_lang/parser"
	"stack_lang/runtime"
	"strconv"
)

func CompileFile(filename string) ([]byte, error) {
    tokens, err := parser.GetTokens(filename)
    fmt.Printf("%v\n", tokens)
    // last := tokens[len(tokens)-1]
    // fmt.Printf("%v\n", []byte(last))
    // tokens = tokens[:len(tokens)-1]
    // fmt.Printf("%v\n", tokens)

    if err != nil && err != io.EOF {
        return nil, fmt.Errorf("CompileFile error: %v", err)
    }

    operations := make([]byte, 0)
    ops := make([]byte, 0)

    for {
        tokens, ops, err = compileTokens(tokens)
        if err != nil {
            return operations, fmt.Errorf("CompileFile: %v", err)
        }

        operations = append(operations, ops...)

        if len(tokens) == 0 {
            break;
        }
    }

    return operations, nil
}

func compileTokens(tokens []string) ([]string, []byte, error) {
    if len(tokens) == 0 {
        return tokens, []byte{}, nil
    }

    first := tokens[0]
    fmt.Println("Token: " + first)

    switch first {
        case "push":
            return compilePush(tokens[1:])
        case "pop":
            return compilePop(tokens[1:])
        case "pops":
            return compilePopStr(tokens[1:])
        case "add":
            return compileAdd(tokens[1:])
        case "sub":
            return compileSub(tokens[1:])
        case "div":
            return compileDiv(tokens[1:])
        case "mul":
            return compileMul(tokens[1:])
        case "print":
            return compilePrintVal(tokens[1:])
        case "prints":
            return compilePrintStr(tokens[1:])
        default:
            return tokens, []byte{}, fmt.Errorf("compileTokens: Unrecognized token: " + first)
    }
}

func compilePop(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.POP}, nil
}

func compilePush(tokens []string) ([]string, []byte, error) {
    if len(tokens) == 0 {
        return tokens, []byte{}, fmt.Errorf("Argument not specified for push")
    }

    if isString(tokens[0]) {
        return compilePushStr(tokens)
    }

    int8Rep, err := strconv.ParseUint(tokens[0], 10, 8)

    newOps := []byte{runtime.PUSH, byte(int8Rep)}

    return tokens[1:], newOps, err
}

func compilePopStr(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.POP_STR}, nil
}

func compilePushStr(tokens []string) ([]string, []byte, error) {
    newOps := make([]byte, 1)

    newOps[0] = runtime.PUSH

    strBytes := []byte(unquoteString(tokens[0]))

    for _,b := range strBytes {
        newOps = append(newOps, b)
    }

    newOps = append(newOps, runtime.END_STR)

    return tokens[1:], newOps, nil
}

func compileAdd(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.ADD}, nil
}

func compileSub(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.SUB}, nil
}

func compileDiv(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.DIV}, nil
}

func compileMul(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.MUL}, nil
}

func compilePrintStr(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.PRINT_STR}, nil
}

func compilePrintVal(tokens []string) ([]string, []byte, error) {
    return tokens, []byte{runtime.PRINT_VAL}, nil
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
