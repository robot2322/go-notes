package service

import "errors"

const StrMaxSize = 1024

var ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

type StringRequest struct {
	A string
	B string
}

type StringService struct {
}

type Service interface {
	Concat(req StringRequest, ret *string) error
}

func (s StringService) Concat(req StringRequest, ret *string) error {
	if len(req.A)+len(req.B) > StrMaxSize {
		*ret = ""
		return ErrMaxSize
	}
	*ret = req.A + req.B
	return nil
}

type ServiceMiddleware func(Service) Service
