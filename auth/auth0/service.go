package auth0

import (
	"github.com/auth0/go-auth0/management"
	"waas-service/auth/jwt/jose"
)

type ManagementAuth0 struct {
	*management.Management
}

func (m *ManagementAuth0) GetUser(userId string) (*management.User, error) {
	user, err := m.User.Read(userId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func NewAuth0Management() *ManagementAuth0 {
	auth0Config := jose.GetAuth0Config()
	m, err := management.New(auth0Config.Domain, management.WithClientCredentials(auth0Config.ClientId, auth0Config.ClientSecret))
	if err != nil {
		return nil
	}
	return &ManagementAuth0{m}
}
