package view

import "github.com/miraikeitai2020/backend-file-proxy/pkg/model/dto"

func NewAppError(code int, err error) dto.Error {
	return dto.Error{Code: code, Message: err.Error()}
}
