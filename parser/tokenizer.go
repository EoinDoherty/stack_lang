package parser

import (
	"bufio"
	"os"
	"unicode"
)

var EOF string = string([]byte{0})

func GetTokens(filename string) ([]string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	return readTokens(reader)
}

func readTokens(reader *bufio.Reader) ([]string, error) {
	tokens := make([]string, 0)

	for {
		s, err := readToken(reader)

		if len(s) > 0 && s != EOF {
			tokens = append(tokens, s)
		}

		if err != nil {
			return tokens, err
		}
	}
}

func readToken(reader *bufio.Reader) (string, error) {
	r, _, err := reader.ReadRune()

	if err != nil {
		return string(r), err
	}

	switch r {
	case '#':
		return "", skipCommentLine(reader)
	case '"':
		return tokenizeString(reader)
	default:
		reader.UnreadRune()
		return tokenizeWord(reader)
	}
}

func skipCommentLine(reader *bufio.Reader) error {
	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			return err
		}

		if r == '\n' {
			return nil
		}
	}
}

func tokenizeString(reader *bufio.Reader) (string, error) {
	str := "\""

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return str, err
		}

		str += string(c)

		if c == '"' {
			return str, nil
		}

	}
}

func tokenizeWord(reader *bufio.Reader) (string, error) {
	word := ""

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			return word, err
		}

		if unicode.IsSpace(r) || r == 0 {
			return word, nil
		}

		word += string(r)
	}
}
