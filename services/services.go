package services

import (
	"errors"
	"fmt"
	"sms/redis"
	"sms/utils"
)

func GenerateOtp(phonenumber string) (string,error) {
	otp := utils.RandomValue(4)
	fmt.Print("code", otp)
	err := redis.SetValue(phonenumber, otp)
	if err != nil {
		return "", err
	}
	token, errs := utils.RandomToken()
	if errs != nil {
		return "", errs
	}

	err = redis.SetValue(token, phonenumber)
	if err != nil {
		return "", err
	}

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