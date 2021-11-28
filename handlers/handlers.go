package handlers

import (
	"encoding/json"
	_ "go/types"
	"net/http"
	"sms/redis"
	"sms/services"
	"sms/utils"
	"sms/zook"
)

type Payload struct {
	PhoneNumber string `json:"phone_number"`
}

type PayloadVerification struct {
	Token string `json:"token"`
	Otp string `json:"otp"`
}

type PayloadVerificationResult struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type Response struct {
	Token string `json:"token"`
}


func PhoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	json.NewDecoder(r.Body).Decode(&payload)
	var phonenumber = payload.PhoneNumber

	value, er := redis.GetValue(phonenumber)
	if value != "" {
		zook.BadRequest(w, "Слишком рано. Повторите попытку позже")
		utils.WriteToSysLog(er)

		return
	}

	token, err := services.GenerateOtp(phonenumber)
	if err != nil {
		zook.BadRequest(w, "Ошибка генерации кода верификации")
		utils.WriteToSysLog(err)

		return
	}
	response := Response{Token: token}
	utils.JsonResponse(w, response)
}

func CodeVerificationHandler(w http.ResponseWriter, r *http.Request) {
	var payload PayloadVerification
	json.NewDecoder(r.Body).Decode(&payload)

	err := services.VerificateOtp(payload.Token, payload.Otp)
	if err != nil {
		zook.BadRequest(w, err)
		return
	}
	response := PayloadVerificationResult{Success: true, Message: "Успешная верификация"}
	utils.JsonResponse(w, response)
}