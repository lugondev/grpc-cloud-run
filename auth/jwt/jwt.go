package jwt

import (
	"context"
	"fmt"
	grpcMetadata "google.golang.org/grpc/metadata"
	"log"
)

type JWT struct {
	validator Validator
}

func New(validator Validator) *JWT {
	return &JWT{
		validator: validator,
	}
}

// Check verifies the jwt token is valid and injects it in the context
func (checker *JWT) Check(metadata grpcMetadata.MD) error {
	// Extract Access Token from context
	if len(metadata.Get("access_token")) == 0 {
		return fmt.Errorf("no accesstoken found in context")
	}

	bearerToken := metadata.Get("access_token")[0]
	if bearerToken == "" {
		return fmt.Errorf("accesstoken is empty")
	}
	log.Println("--> bearer token: ", bearerToken)
	// Parse and validate token injected in context
	claims, err := checker.validator.ValidateToken(context.Background(), bearerToken)
	if err != nil {
		return err
	}

	log.Println("--> user claims: ", claims)
	return nil
}
