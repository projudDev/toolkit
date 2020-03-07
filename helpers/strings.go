package helpers

import (
	"fmt"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func StrReturnNumberOnly(s string) string {

	result := ""

	for _, c := range s {
		if unicode.IsDigit(c) {
			result = result + string(c)
		}
	}

	return result

}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func RemoveAcentos(texto string) string {

	b := make([]byte, len(texto))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, e := t.Transform(b, []byte(texto), true)
	if e != nil {
		panic(e)
	}

	fmt.Println(string(b))
	return string(b)
}
