package apperror

import (
	"context"
	"errors"
	"upserv/config"
)

var (
	BadRequestGeneral = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Error on parse from-data",
			Code:      1,
			HttpCode:  400,
			RequestID: requestId,
		}
	}
	BadRequestOnParams = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Require field is not set",
			Code:      2,
			HttpCode:  400,
			RequestID: requestId,
		}
	}
	AuthValidation = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Error on auth validation",
			Code:      3,
			HttpCode:  401,
			RequestID: requestId,
		}
	}
	NotFound = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Entity not found",
			Code:      4,
			HttpCode:  404,
			RequestID: requestId,
		}
	}
	InvalidCode = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Code is not valid",
			Code:      5,
			HttpCode:  404,
			RequestID: requestId,
		}
	}
	Unauthorized = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "User unauthorized",
			Code:      6,
			HttpCode:  401,
			RequestID: requestId,
		}
	}
	Forbidden = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Forbidden request",
			Code:      7,
			HttpCode:  403,
			RequestID: requestId,
		}
	}
	DBError = func(ctx context.Context, err error) *IError {
		if err == nil {
			err = errors.New("an error occurred while executing query")
		}
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       err.Error(),
			Code:      8,
			HttpCode:  500,
			RequestID: requestId,
		}
	}
	UploadFailed = func(ctx context.Context, err error) *IError {
		if err == nil {
			err = errors.New("an error occurred while uploading file")
		}
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       err.Error(),
			Code:      9,
			HttpCode:  500,
			RequestID: requestId,
		}
	}
	FileNotFound = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "File not found",
			Code:      10,
			HttpCode:  404,
			RequestID: requestId,
		}
	}
	InternalError = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Internal server error",
			Code:      11,
			HttpCode:  500,
			RequestID: requestId,
		}
	}
	ServiceError = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "External service error",
			Code:      12,
			HttpCode:  500,
			RequestID: requestId,
		}
	}
	ForbiddenMessage = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Feeling lonely? You can't talk to yourself.",
			Code:      13,
			HttpCode:  403,
			RequestID: requestId,
		}
	}
	ForbiddenLogin = func(ctx context.Context) *IError {
		requestId, _ := ctx.Value(config.RequestIdKey).(string)
		return &IError{
			Msg:       "Wrong login or password!",
			Code:      14,
			HttpCode:  403,
			RequestID: requestId,
		}
	}
)

// Error struct
type IError struct {
	Msg       string
	Code      int32
	RequestID string
	HttpCode  int
}

func (e *IError) WithMsg(msg string) *IError {
	e.Msg = msg
	return e
}

func (e *IError) Error() string {
	return e.Msg
}
