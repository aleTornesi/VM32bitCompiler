package tokenizer

import (
	"regexp"
)

type Token struct {
    Type string
    Value string
}

func TokenizeSingleChar(input string, current uint, char byte, _type string) (bool, uint, *Token) {
    if current >= uint(len(input)) {
        panic("Invalid cursor")
    }
    if input[current] == char {
        return true, 1, &Token{_type, string(char)}
    }
    return false, 0, nil
}

func TokenizePattern(input string, current uint, firstCharRegex, restOfStringRegex string, Type string) (bool, uint, *Token) {
    if current >= uint(len(input)) {
        panic("Invalid cursor")
    }
    var consumedChars uint = 0
    isValidCharacter, _ := regexp.MatchString(firstCharRegex, string(input[current]))
    if isValidCharacter {
        hasOverflowed := false
        for current + consumedChars < uint(len(input)) && isValidCharacter {
            consumedChars++
            if current + consumedChars >= uint(len(input)) {
                hasOverflowed = true
                break
            }
            isValidCharacter, _ = regexp.MatchString(restOfStringRegex, string(input[current + consumedChars]))
        }
        var value string
        if hasOverflowed {
            value = input[current:]
        } else {
            value = input[current: current + consumedChars]
        }
        return true, consumedChars, &Token{Type, value}
    }
    return false, 0, nil
}

func TokenizeString(input string, current uint) (bool, uint, *Token) {
    if current >= uint(len(input)) {
        panic("Invalid cursor")
    }
    if input[current] == '"' {
        var consumedChars uint = 1
        if current + consumedChars >= uint(len(input)) {
            panic("Unterminated string")
        }
        for input[current + consumedChars] != '"' {
            if input[current + consumedChars] == '\\' {
                consumedChars++
            }
            consumedChars++
            if current + consumedChars >= uint(len(input)) {
                panic("Unterminated string")
            }
        }
        consumedChars++
        return true, consumedChars, &Token{"string", input[current: current + consumedChars]}
    }
    return false, 0, nil
}
var x int = 0
var tokenizers []func(string, uint) (bool, uint, *Token) = []func(string, uint) (bool, uint, *Token) {
    func(input string, cursor uint) (bool, uint, *Token) {return TokenizePattern(input, cursor, "[[:space:]]", "[[:space:]]", "whitespace")},
    func(input string, cursor uint) (bool, uint, *Token) {return TokenizeSingleChar(input, cursor, '(', "lparen")},
    func(input string, cursor uint) (bool, uint, *Token) {return TokenizeSingleChar(input, cursor, ')', "rparen")},
    func(input string, cursor uint) (bool, uint, *Token) {return TokenizePattern(input, cursor, "[0-9]", "[0-9]", "number")},
    func(input string, cursor uint) (bool, uint, *Token) {return TokenizePattern(input, cursor, "[A-z_]", "[A-z0-9_]", "name")},
    TokenizeString,
}

func Tokenize(input string) []Token {
    var current uint = 0
    tokens := []Token{}
    for current < uint(len(input)) {
        for _, t := range(tokenizers) {
            isRightPattern, chars, token := t(input, current)
            if isRightPattern {
                current += chars
                if token.Type != "whitespace" {
                    tokens = append(tokens, *token)
                }
                break
            }
        }
    }
    return tokens
}
