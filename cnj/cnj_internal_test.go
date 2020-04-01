package cnj

import "testing"
import "fmt"

func TestCnj(t *testing.T) {
	err := ValidaCNJ("12334234242342342343")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("Ok")
}
