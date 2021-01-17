package parser

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestReadTokens(t *testing.T) {
	validateTokenizer(t, "push 1", []string{"push", "1"})
	validateTokenizer(t, "push 1 asdf", []string{"push", "1", "asdf"})
	validateTokenizer(t, "#asdf qwe", []string{})

	multiline := `push 1 
	# comment
	thing1 thing2 #comment
	1`

	expected := []string{"push", "1", "thing1", "thing2", "1"}

	validateTokenizer(t, multiline, expected)
}

func TestReadToken(t *testing.T) {
	validateSingleToken(t, readToken, " asdf qwe", "")
	validateSingleToken(t, readToken, "asdf qwe", "asdf")
	validateSingleToken(t, readToken, "5 qwe", "5")
	validateSingleToken(t, readToken, "\"word1 word2\" 123", "\"word1 word2\"")
	validateSingleToken(t, readToken, "# invisible :O", "")

	validateSingleToken(t, tokenizeString, "asdf asdf\"", "\"asdf asdf\"")
	validateSingleToken(t, tokenizeString, "\"", "\"\"")

	validateSingleToken(t, tokenizeWord, "asdf qwe", "asdf")
	validateSingleToken(t, tokenizeWord, " asdf qwe", "")
	validateSingleToken(t, tokenizeWord, "a", "a")
}

func validateTokenizer(t *testing.T,
	input string,
	expected []string) {
	stringReader := bufio.NewReader(strings.NewReader(input))
	tokens, err := readTokens(stringReader)

	if err != nil && err != io.EOF {
		t.Errorf("Tokenizing error: %v", err)
	}

	for i, token := range tokens {
		expectedToken := expected[i]
		if token != expectedToken {
			t.Errorf("Tokenizing mismatch: %s != %s", token, expectedToken)
		}
	}

	if len(expected) != len(tokens) {
		t.Errorf("Number of tokens does not match expected number of tokens")
	}
}

func validateSingleToken(
	t *testing.T,
	f func(*bufio.Reader) (string, error),
	input string,
	expected string) {

	stringReader := bufio.NewReader(strings.NewReader(input))

	output, err := f(stringReader)

	if err != nil && err != io.EOF {
		t.Errorf("Tokenizing error: %v", err)
	}

	if output != expected {
		t.Errorf("Tokenizing error: %s != %s", output, expected)
	}
}
