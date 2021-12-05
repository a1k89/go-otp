package handlers

import (
	"encoding/json"
	_ "go/types"
	"net/http"
	"sms/models"
	"sms/services"
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
	// Generate OTP
	token, err := services.GenerateOtp(payload.PhoneNumber)
	if err != nil {
		zook.BadRequest(w, "Ошибка генерации кода верификации")
		return
	}
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