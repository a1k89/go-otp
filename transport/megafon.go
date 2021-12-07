package transport

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Megafon struct {
	To string
	From string
	Message string
}

func (t Megafon) hash() string {
	login := os.Getenv("TRANSPORT_CRED_LOGIN")
	password := os.Getenv("TRANSPORT_CRED_PASSWORD")
	encoding := login + ":" + password
	hashValue := base64.StdEncoding.EncodeToString([]byte(encoding))

	return hashValue

}
func (t Megafon) SendSms() {
	jsonData, _ := json.Marshal(t)
	header := fmt.Sprintf("Basic %s", t.hash())
	request, err := http.NewRequest("POST", os.Getenv("TRANSPORT_CRED_URL"),bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("NewRequest error: %s", err)
		return
	}
	request.Header.Set("Authorization", header)
	request.Header.Set("content-type","application/json")
	client := http.Client{}

	_, err = client.Do(request)
	if err != nil {
		fmt.Printf("Send message error: %s", err)
	}

}