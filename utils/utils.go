package utils

import (
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"log/syslog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RandomValue(size int) string {
	otp := os.Getenv("OTP")
	if otp != "" {
		return otp
	}
	var result string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		rdm := rand.Intn(9) + 1
		result += strconv.Itoa(rdm)
	}

	return result
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func RandomToken() (string, error) {
	u4, err := uuid.NewV4()
	UUIDtoken := u4.String()
	if err != nil {
		return "", err
	}

	return UUIDtoken, nil
}

func WriteToSysLog(err error) {
	logWriter, er := syslog.New(syslog.LOG_SYSLOG, "SMS")
	if er != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	log.SetFlags(log.Lshortfile)
	log.SetOutput(logWriter)
	log.Println(err)
}