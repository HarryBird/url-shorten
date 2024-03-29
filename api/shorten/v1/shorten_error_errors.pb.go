// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsCreatrShortenUrlFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CREATR_SHORTEN_URL_FAIL.String() && e.Code == 500
}

func ErrorCreatrShortenUrlFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_CREATR_SHORTEN_URL_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsGetShortenUrlFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_SHORTEN_URL_FAIL.String() && e.Code == 500
}

func ErrorGetShortenUrlFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GET_SHORTEN_URL_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsDeleteShortenUrlFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DELETE_SHORTEN_URL_FAIL.String() && e.Code == 500
}

func ErrorDeleteShortenUrlFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DELETE_SHORTEN_URL_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsDecodeShortenUrlFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DECODE_SHORTEN_URL_FAIL.String() && e.Code == 500
}

func ErrorDecodeShortenUrlFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DECODE_SHORTEN_URL_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsShortenUrlNonexist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SHORTEN_URL_NONEXIST.String() && e.Code == 400
}

func ErrorShortenUrlNonexist(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_SHORTEN_URL_NONEXIST.String(), fmt.Sprintf(format, args...))
}

func IsShortenIdInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SHORTEN_ID_INVALID.String() && e.Code == 400
}

func ErrorShortenIdInvalid(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_SHORTEN_ID_INVALID.String(), fmt.Sprintf(format, args...))
}

func IsShortenCodeInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SHORTEN_CODE_INVALID.String() && e.Code == 400
}

func ErrorShortenCodeInvalid(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_SHORTEN_CODE_INVALID.String(), fmt.Sprintf(format, args...))
}
