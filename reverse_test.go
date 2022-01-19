package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	testString string
	expString  string
}

func TestReverse(t *testing.T) {
	tests := map[string]test{
		"simple string":          {testString: "Hello", expString: "olleH"},
		"empty string":           {testString: "", expString: ""},
		"one letter string":      {testString: "a", expString: "a"},
		"two letter string":      {testString: "no", expString: "on"},
		"string with whitespace": {testString: "\t whitespace   ", expString: "   ecapsetihw \t"},
		"cyrillic string":        {testString: "Привет", expString: "тевирП"},
		"string with emojis":     {testString: "I like tennis 🤚🎾😀", expString: "😀🎾🤚 sinnet ekil I"},
		"haiku": {
			testString: `My code does not work
Seventh hour of debugging
Life is miserable`,
			expString: `elbaresim si efiL
gniggubed fo ruoh htneveS
krow ton seod edoc yM`,
		},
	}

	for _, tt := range tests {
		res := ReverseString(tt.testString)
		assert.Equal(t, tt.expString, res)
	}
}
