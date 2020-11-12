package dto

type ConfigUpdateRequest struct {
	URL       string `json:"url"`
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
}
