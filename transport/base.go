package transport

import (
	"fmt"
	"time"
)

type transport interface {
	SendSms() bool
}

type Transport struct {
	To string
	From string
	Message string
}

func (t *Transport) SendSms() bool {
	time.Sleep(10 * time.Second)
	fmt.Print("Transport send SMS", t.Message)
	return true
}