package transport

import (
	"fmt"
	"time"
)

type Megafon struct {
	To string
	From string
	Message string
}

func (t Megafon) SendSms() {
	time.Sleep(10 * time.Second)
	fmt.Print("Transport send SMS", t.Message)
}