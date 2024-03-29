package biz

import "errors"

var (
	ErrParamInvalid      = errors.New("invalid param")
	ErrNotFoundFromDB    = errors.New("data not found from db")
	ErrNotFoundFromRedis = errors.New("data not found from redis")
	ErrURLCodeNonexist   = errors.New("shorten url code not exist")
	ErrTimeout           = errors.New("timeout")
)
