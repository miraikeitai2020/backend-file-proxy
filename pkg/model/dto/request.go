package dto

type ConfigUpdateRequest struct {
	URL       string `json:"url"`
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
}

type CreateImageRequest struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}
