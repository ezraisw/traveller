package traveller

import (
	"errors"
	"strings"
)

var (
	// The error that is returned when there is an invalid path.
	ErrInvalidPath = errors.New("invalid path")
)

// Convert a string path to a series of matchers.
//
// Will panic if the path is invalid.
func MustPath(ps string, caseInsensitive bool) []Matcher {
	mp, err := Path(ps, caseInsensitive)
	if err != nil {
		panic(err)
	}
	return mp
}

// Convert a string path to a series of matchers.
func Path(ps string, caseInsensitive bool) ([]Matcher, error) {
	tokens := splitEscape(ps, '.', '\\')
	matchers := make([]Matcher, 0, len(tokens))
	for _, token := range tokens {
		if isExactToken(token) {
			if !caseInsensitive {
				matchers = append(matchers, MatchExact{Value: token})
			} else {
				matchers = append(matchers, MatchPattern{
					Pattern: token,
					Options: MatchPatternOptions{CaseInsensitive: caseInsensitive},
				})
			}
		} else if isMultiMatchToken(token) {
			matchers = append(matchers, MatchMulti{})
		} else if !isInvalidToken(token) {
			matchers = append(matchers, MatchPattern{
				Pattern: token,
				Options: MatchPatternOptions{CaseInsensitive: caseInsensitive},
			})
		} else {
			return nil, ErrInvalidPath
		}
	}
	return matchers, nil
}

// Whether the token is an exact match value.
func isExactToken(token string) bool {
	return !strings.Contains(token, "*")
}

// Whether the token is a multi match/recursive match value.
func isMultiMatchToken(token string) bool {
	return token == "**"
}

// Whether a token is invalid and should not be parsed.
func isInvalidToken(token string) bool {
	return strings.Contains(token, "**") && len(token) != 2
}

// Splits a string to a collection of token by the given separator.
//
// Will not attempt to split when a separator is preceded by
// the specified escape character.
func splitEscape(s string, separator, escape byte) []string {
	var (
		token  []byte
		tokens []string
	)
	for i := 0; i < len(s); i++ {
		if s[i] == separator {
			tokens = append(tokens, string(token))
			token = token[:0]
		} else if s[i] == escape && i+1 < len(s) {
			i++
			token = append(token, s[i])
		} else {
			token = append(token, s[i])
		}
	}
	tokens = append(tokens, string(token))
	return tokens
}
