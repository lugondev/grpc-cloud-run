package jwt

import (
	"context"
	"waas-service/dto"
)

type Validator interface {
	ValidateToken(ctx context.Context, token string) (*dto.UserClaims, error)
}
