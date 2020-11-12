package view

import "github.com/miraikeitai2020/backend-file-proxy/pkg/model/dto"

func NewAppError(code int, err error) dto.Error {
	return dto.Error{Code: code, Message: err.Error()}
}

func NewMinioAccessInfo(url, pk, sk string) dto.AccessInfo {
	return dto.AccessInfo{Endpoint: url, PublicKey: pk, SecretKey: sk}
}

func NewConfigUpdateResponse(info *dto.AccessInfo, err *dto.Error) dto.ConfigUpdateResponse {
	return dto.ConfigUpdateResponse{Info: info, Err: err}
}
