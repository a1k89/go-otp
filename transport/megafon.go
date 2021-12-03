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
	request, err := http.Post(os.Getenv("TRANSPORT_CRED_URL"), "application/json", bytes.NewBuffer(jsonData))
	request.Header.Set("Authorization", header)
	if err != nil {
		fmt.Print("Error", err)
	}
	fmt.Print("Request", request)
}