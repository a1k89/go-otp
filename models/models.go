package models

import (
	"errors"
	"regexp"
)

type Payload struct {
	PhoneNumber string `json:"phone_number"`
}

func (p *Payload) IsValid() error {
	pattern := regexp.MustCompile(`\d{11,20}`)
	matched := pattern.MatchString(p.PhoneNumber)
	if !matched {
		return errors.New("Номер телефона должен содержать только цифры")
	}

	return nil
}

type PayloadResult struct {
	Token string `json:"token"`
}

type PayloadVerification struct {
	Token string `json:"token"`
	Otp string `json:"otp"`
}

type PayloadVerificationResult struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

