package dto

type ConfigUpdateResponse struct {
	Info *AccessInfo `json:"accessConfig,omitempty"`
	Err  *Error      `json:"error,omitempty"`
}

type AccessInfo struct {
	Endpoint  string `json:"endpoint"`
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
}

type ObjectInfo struct {
	ID   string `json:"id"`
	Size int64  `json:"size"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
