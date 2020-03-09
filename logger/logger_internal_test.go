package logger

import "testing"
import "errors"
func TestLogger(t *testing.T) {
	var err error
	err = errors.New("New test error")
	ProductError(int64(1), err)
}