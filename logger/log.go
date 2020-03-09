package logger

import (
	"fmt"
	"github.com/projudDev/toolkit/net"
	"log"
	"runtime"
	"crypto/tls"
	"net/http"
)

func Error(err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
}

func ProductError(productID int64, err error) (b bool) {
	if err != nil {
		ip, _ := net.ExternalIP()
		pc, fn, line, _ := runtime.Caller(1)
		description := fmt.Sprintf("at=[error] product=[%v] ip[%s] module=[%s] path=[%s:%d] description=[%v]", productID, ip, runtime.FuncForPC(pc).Name(), fn, line, err)
		log.Printf(description)
		SendToTelegram(description)
		b = true
	}
	return
}

func SendToTelegram(description string) {
	URL := "https://api.telegram.org/bot708536877:AAEtR-PxVPhiH1OhLj6w43O2EFXuDiGespg/sendMessage?chat_id=-288258884&text=" + description
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", URL, nil)
	client.Do(req)
}
