package view

import "github.com/miraikeitai2020/backend-file-proxy/pkg/model/dto"

func NewAppError(code int, err error) dto.Error {
	return dto.Error{Code: code, Message: err.Error()}
}

func NewObjectInfo(id string, size int64) dto.ObjectInfo {
	return dto.ObjectInfo{ID: id, Size: size}
}
