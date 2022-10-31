package jose

import (
	"context"
	"encoding/json"
	"errors"
	"waas-service/dto"
)

type CustomClaims struct {
	UserClaims    *dto.UserClaims
	userClaimPath string
}

func NewCustomClaims(path string) *CustomClaims {
	return &CustomClaims{
		userClaimPath: path,
	}
}

func (c *CustomClaims) UnmarshalJSON(data []byte) error {
	var res map[string]interface{}
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	if _, ok := res[c.userClaimPath]; ok {
		c.UserClaims = &dto.UserClaims{}
		bClaims, _ := json.Marshal(res[c.userClaimPath])
		if err := json.Unmarshal(bClaims, &c.UserClaims); err != nil {
			return errors.New("invalid custom claims format")
		}
	} else {
		return errors.New("missing custom claims data")
	}

	return nil
}

func (c *CustomClaims) Validate(_ context.Context) error {
	return nil
}
