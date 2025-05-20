package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

var ErrRecordNotFound = errors.New("record not found")

func ErrCannotGetEntity(entity string, err error) *AppError {
	msg := "Can not get " + entity
	return NewCustomError(err, msg, "ErrCannotGetItem")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	msg := "Can not create " + entity
	return NewCustomError(err, msg, "ErrCannotGetItem")
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	msg := "Can not update " + entity
	return NewCustomError(err, msg, "ErrCannotGetItem")
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	msg := "Can not delete " + entity
	return NewCustomError(err, msg, "ErrCannotGetItem")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Internal Server Error", err.Error(), "ErrInternal")
}
