package email

import "testing"
import "github.com/projudDev/toolkit/driver"
import "fmt"

func TestSmtpServer(t *testing.T) {
	db, err := driver.NewConnectionMysql()
	if err != nil {
		t.Error(err.Error())
	}
	s, err := NewSmtpServer(db.SQL, 131)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(s)
}
