package services

import (
	"errors"
	"sms/redis"
	"sms/transport"
	"sms/utils"
)

func GenerateOtp(phonenumber string) (string, error) {
	// Check if not exist
	value, _ := redis.GetValue(phonenumber)
	if value != "" {
		return "", errors.New("Слишком рано. Повторите попытку позже")
	}
	// Generate and set random value
	otp := utils.RandomValue(4)
	err := redis.SetValue(phonenumber, otp)
	if err != nil {
		return "", err
	}
	// Generate and set token
	token, errs := utils.RandomToken()
	if errs != nil {
		return "", errs
	}

	err = redis.SetValue(token, phonenumber)
	if err != nil {
		return "", err
	}

	// Send OTP to user mobile phone
	mgf := transport.Megafon{
		To: phonenumber,
		From: phonenumber,
		Message: otp}

	go transport.SendMessage(mgf)

	return token, nil
}

func VerificateOtp(token string, otp string) error {
	phonenumber, err := redis.GetValue(token)
	if err != nil {
		return errors.New("Неверный токен или срок действия его истек")
	}

	redisOtp, err := redis.GetValue(phonenumber)
	if err != nil {
		return errors.New("Код не найден")
	}

	if redisOtp != otp {
		return errors.New("Неверный код верификации")
	}

	return nil
}