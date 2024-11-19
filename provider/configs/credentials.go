package configs

type IntegrationCredentials struct {
	Token     *string `json:"token"`
	Email     *string `json:"email"`
	APIKey    *string `json:"api_key"`
	AccessKey *string `json:"access_key"`
	SecretKey *string `json:"secret_key"`
}
