package proto

import (
	"context"
	"sms/models"
	"sms/proto/github.com/monkrus/grpc-from0"
	"sms/services"
)

type Server struct {}

func (Server) Generate(ctx context.Context, request *grpc_from0.PayloadGenerateRequest) (*grpc_from0.PayloadGenerateResponse, error) {
	payload := models.Payload{PhoneNumber: request.PhoneNumber}
	err := payload.IsValid()
	if err != nil {
		return nil, err
	}
	token, err := services.GenerateOtp(payload.PhoneNumber)
	if err != nil {
		return nil, err
	}
	response := grpc_from0.PayloadGenerateResponse{Token: token}

	return &response, nil
}

func (Server) Verificate(ctx context.Context, request *grpc_from0.PayloadVerificateRequest) (*grpc_from0.PayloadVerificateResponse, error) {
	err := services.VerificateOtp(request.Token, request.Otp)
	if err != nil {
		return nil, err
	}

	var response = grpc_from0.PayloadVerificateResponse{Message: "Успешная верификация по коду", Success: true}

	return &response, nil
}