package jose

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/gommon/log"
	"waas-service/auth/utils"
	"waas-service/dto"
)

type Validator struct {
	validator *validator.Validator
}

func NewValidator() (*Validator, error) {
	cfg := GetOIDCConfig()
	issuerURL, err := url.Parse(cfg.Issuer)
	if err != nil {
		return nil, err
	}
	log.Info("OIDC config:", cfg)
	v, err := validator.New(
		jwks.NewCachingProvider(issuerURL, cfg.CacheTTL).KeyFunc,
		validator.RS256,
		issuerURL.String(),
		cfg.Audience,
		validator.WithCustomClaims(func() validator.CustomClaims {
			return NewCustomClaims(cfg.ClaimsPath)
		}),
	)
	if err != nil {
		return nil, err
	}

	return &Validator{validator: v}, nil
}

func (v *Validator) ValidateToken(ctx context.Context, token string) (*dto.UserClaims, error) {

	userCtx, err := v.validator.ValidateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	claims := userCtx.(*validator.ValidatedClaims)
	log.Printf("Claims: %v", claims)
	if claims.CustomClaims != nil {
		if orchestrateUserClaims := claims.CustomClaims.(*CustomClaims).UserClaims; orchestrateUserClaims != nil {
			return orchestrateUserClaims, nil
		}

		return nil, fmt.Errorf("expected custom claims not found")
	}

	// The tenant ID is the "sub" field, then is "tenant_id:username" or "tenant_id"
	claim := &dto.UserClaims{}
	sub := claims.RegisteredClaims.Subject
	pieces := strings.Split(sub, utils.AuthSeparator)
	if len(pieces) == 0 {
		claim.TenantId = pieces[0]
	} else {
		claim.Username = pieces[len(pieces)-1]
		claim.TenantId = strings.Replace(sub, utils.AuthSeparator+claim.Username, "", 1)
	}

	return claim, nil
}
