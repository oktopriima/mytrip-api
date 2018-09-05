package auth

type tokenResponse struct {
	AccessToken  string      `json:"access_token"`
	RefrashToken string      `json:"refresh_token"`
	TokenType    string      `json:"type_type"`
	ExpiresIn    float64     `json:"expires_in"`
	ExpiredAt    int64       `json:"expired_at"`
	Data         interface{} `json:"data"`
}
