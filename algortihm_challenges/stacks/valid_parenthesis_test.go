package stacks

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func Test_isValid(t *testing.T) {
	const s1 = "()"
	assert.True(t, isValid(s1))
	const s2 = "()[]{}"
	assert.True(t, isValid(s2))
	const s3 = "(]"
	assert.False(t, isValid(s3))
}

func Test_isValidLevelUp(t *testing.T) {
	const s1 = "{[]}"
	assert.True(t, isValid(s1))
	const s2 = "()[]{}"
	assert.True(t, isValid(s2))
	const s3 = "(]"
	assert.False(t, isValid(s3))
}

func Test_IsValidExtensiveTesting(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single opening parenthesis",
			s:    "(",
			want: false,
		},
		{
			name: "single closing parenthesis",
			s:    ")",
			want: false,
		},
		{
			name: "single opening brace",
			s:    "{",
			want: false,
		},
		{
			name: "single closing brace",
			s:    "}",
			want: false,
		},
		{
			name: "single opening bracket",
			s:    "[",
			want: false,
		},
		{
			name: "single closing bracket",
			s:    "]",
			want: false,
		},
		{
			name: "two opening parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "two closing parentheses",
			s:    "))",
			want: false,
		},
		{
			name: "opening and closing parenthesis",
			s:    "( )",
			want: true,
		},
		{
			name: "closing and opening parenthesis",
			s:    ") (",
			want: false,
		},
		{
			name: "two opening braces",
			s:    "{}",
			want: true,
		},
		{
			name: "two closing braces",
			s:    "}}",
			want: false,
		},
		{
			name: "opening and closing brace",
			s:    "{ }",
			want: true,
		},
		{
			name: "closing and opening brace",
			s:    "} {",
			want: false,
		},
		{
			name: "two opening brackets",
			s:    "[]",
			want: true,
		},
		{
			name: "two closing brackets",
			s:    "]]",
			want: false,
		},
		{
			name: "opening and closing bracket",
			s:    "[ ]",
			want: true,
		},
		{
			name: "closing and opening bracket",
			s:    "] [",
			want: false,
		},
		{
			name: "parenthesis inside braces",
			s:    "{ ( ) }",
			want: true,
		},
		{
			name: "brackets inside braces",
			s:    "{ [ ] }",
			want: true,
		},
		{
			name: "braces inside brackets",
			s:    "[ { } ]",
			want: true,
		},
		{
			name: "nested parentheses",
			s:    "(( ))",
			want: true,
		},
		{
			name: "nested braces",
			s:    "{{ }}",
			want: true,
		},
		{
			name: "nested brackets",
			s:    "[[ ]]",
			want: true,
		},
		{
			name: "nested different brackets",
			s:    "([ { } ])",
			want: true,
		},
		{
			name: "opening parenthesis followed by opening brace",
			s:    "( {",
			want: false,
		},
		{
			name: "opening brace followed by opening parenthesis",
			s:    "{ (",
			want: false,
		},
		{
			name: "opening brace followed by closing parenthesis",
			s:    "{ )",
			want: false,
		},
	}
	for _, test := range testCases {
		log.Default().Println(test.name)
		assert.Equal(t, test.want, isValid(test.s))
	}
}

func Test_IsValidHighDifficulty(t *testing.T) {
	//First 3333 characters are open brackets in the order of ([{, and the next 3333 characters are close brackets in the order of )]}.
	//This ensures that every open bracket is followed by a different type of close bracket.
	inputFalse := strings.Repeat("([{", 3333) + strings.Repeat(")]}", 3333)
	expectedFalse := false
	assert.Equal(t, expectedFalse, isValid(inputFalse))

	inputTrue := strings.Repeat("{[(", 3333) + strings.Repeat(")]}", 3333)
	expectedTrue := true
	assert.Equal(t, expectedTrue, isValid(inputTrue))
}
