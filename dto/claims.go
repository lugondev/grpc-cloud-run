package dto

// UserClaims represent raw claims extracted from an authentication method
type UserClaims struct {
	TenantId string `json:"tenant_id"`
	Username string `json:"username"`

	EmailVerified bool     `json:"email_verified"`
	Roles         []string `json:"roles"`
}
