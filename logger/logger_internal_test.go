package logger

import "testing"
import "errors"

func TestLogger(t *testing.T) {
	var err error
	err = errors.New("New test error")
	ProductError(int64(1), err)
}

const markdow = `<pre><code>pre-formatted fixed-width code block written in the Go programminglanguage</code></pre>`

func TestHtmlError(t *testing.T) {
	var err error
	err = errors.New(string(markdow))
	MarkDowError(err)
}
