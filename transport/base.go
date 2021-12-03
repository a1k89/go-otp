package transport


type Transport interface {
	SendSms()
}

func SendMessage(t Transport) {
	t.SendSms()
}