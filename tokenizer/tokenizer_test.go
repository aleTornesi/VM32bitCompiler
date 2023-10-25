package tokenizer

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTokenizeSingleChar(t *testing.T) {
    isValid, consumedChars, token := TokenizeSingleChar("(", 0, '(', "lparen")
    if !isValid || consumedChars != 1 || !cmp.Equal (*token, Token{"lparen", "("})  {
        t.Fatalf("tokenizeSingleChar(\"(\", 0, ')') = %v, %v, %v; expected true, 1, Token{\"lparen\", \"(\"})", isValid, consumedChars, *token)
    }
}

func TestPatterns(t *testing.T) {
    _, err := regexp.MatchString("[A-z_]", "temp")
    if err != nil {
        t.Fatalf("[A-z_] must be compiled")
    }
}

func TestTokenizePattern(t *testing.T) {
    isValid, consumedChars, token := TokenizePattern("temp", 0, "[A-z_]", "[A-z0-9_]", "name")
    if !isValid || consumedChars != 4 || !cmp.Equal (*token, Token{"name", "temp"})  {
        t.Fatalf("TokenizePattern(\"temp\", 0, \"^[A-z][A-z0-9_]*$\", \"name\") = %v, %v, %v; expected true, 4, Token{\"name\", \"temp\"})", isValid, consumedChars, *token)
    }
}
func TestTokenizeString(t *testing.T) {
    isValid, consumedChars, token := TokenizeString("\"temp\"", 0)
    if !isValid || consumedChars != 6 || !cmp.Equal (*token, Token{"string", "\"temp\""})  {
        t.Fatalf("TokenizeString(\"\"temp\"\", 0) = %v, %v, %v; expected true, 1, Token{\"string\", \"\"temp\"\"})", isValid, consumedChars, *token)
    }
}

func TestTokenize(t *testing.T) {
    tokens := Tokenize("sampleFunction(\"parameter\")")
    expectedTokens := []Token{Token{"name", "sampleFunction"}, Token{"lparen", "("}, Token{"string", "\"parameter\""}, Token{"rparen", ")"}}
    if !cmp.Equal(tokens, []Token{Token{"name", "sampleFunction"}, Token{"lparen", "("}, Token{"string", "\"parameter\""}, Token{"rparen", ")"}}) {
        t.Fatalf("tokenize(\"sampleFunction(\"parameter\")\") = %v; expected %v", tokens, expectedTokens)
    }
}

