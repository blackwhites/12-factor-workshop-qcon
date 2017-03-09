package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetDealEndpoint   endpoint.Endpoint
	GetSecretEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetDealEndpoint:   MakeGetDealEndpoint(s),
		GetSecretEndpoint: MakeGetSecretEndpoint(s),
	}
}

func MakeGetDealEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDealRequest)
		p, e := s.GetDeal(req.ID)
		return getDealResponse{Id: p.Id, Name: p.Name}, e
	}
}

func MakeGetSecretEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getSecretRequest)
		message, e := s.GetSecret(req.Code)
		return getSecretResponse{Message: message}, e
	}
}

type getDealRequest struct {
	ID int
}

type getDealResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type getSecretRequest struct {
	Code string
}

type getSecretResponse struct {
	Message string `json:"message"`
}
