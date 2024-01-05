package iteration

import (
	"strings"
)

func Repeat(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}

func StringsRepeat(character string, repeatCount int) string {
	repeated := strings.Repeat(character, repeatCount)
	return repeated
}

func StringsBuilderRepeat(character string, repeatCount int) string {
	var b strings.Builder
	for i := 0; i < repeatCount; i++ {
		b.WriteString(character)
	}
	return b.String()
}

func StringsBuilderWithRunesRepeat(character string, repeatCount int) string {
	var b strings.Builder
	c := []rune(character)
	for i := 0; i < repeatCount; i++ {
		for _, r := range c {
			b.WriteRune(r)
		}
	}
	return b.String()
}
