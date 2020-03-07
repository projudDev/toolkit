package helpers

import "testing"
import "github.com/projudDev/toolkit/driver"
import "fmt"
func TestSmtpServer(t *testing.T) {
	db, err := driver.NewConnectioMysql()
	if err != nil {
		t.Error(err.Error())
	}
	s, err := NewSmtpServer(db.SQL, 10)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(s)
}
