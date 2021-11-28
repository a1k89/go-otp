package handlers

import (
	"encoding/json"
	_ "go/types"
	"net/http"
	"sms/models"
	"sms/redis"
	"sms/services"
	"sms/transport"
	"sms/utils"
	"sms/zook"
)

func PhoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Get phone number and validate it
	var payload models.Payload
	json.NewDecoder(r.Body).Decode(&payload)
	err := payload.IsValid()
	if err != nil {
		zook.BadRequest(w, err)
		return
	}

	// Check availability to create OTP
	var phonenumber = payload.PhoneNumber
	value, er := redis.GetValue(phonenumber)
	if value != "" {
		zook.BadRequest(w, "Слишком рано. Повторите попытку позже")
		utils.WriteToSysLog(er)
		return
	}

	// Generate OTP
	token,otp, err := services.GenerateOtp(phonenumber)
	if err != nil {
		zook.BadRequest(w, "Ошибка генерации кода верификации")
		utils.WriteToSysLog(err)
		return
	}

	// Send OTP to user mobile phone
	transportHandler := transport.Transport{
		To: phonenumber,
		From: phonenumber,
		Message: otp }
	go transportHandler.SendSms()

	response := models.PayloadResult{Token: token}
	utils.JsonResponse(w, response)
}

func CodeVerificationHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.PayloadVerification
	json.NewDecoder(r.Body).Decode(&payload)

	err := services.VerificateOtp(payload.Token, payload.Otp)
	if err != nil {
		zook.BadRequest(w, err)
		return
	}
	response := models.PayloadVerificationResult{
		Success: true,
		Message: "Успешная верификация"}
	utils.JsonResponse(w, response)
}