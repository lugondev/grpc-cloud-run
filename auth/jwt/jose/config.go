package jose

import (
	"time"

	"github.com/spf13/viper"
)

var (
	appConfig *AppConfig
)

const (
	issuerKey     = "auth.jwt.issuer"
	issuerEnv     = "AUTH_ISSUER"
	audienceKey   = "auth.jwt.audience"
	audienceEnv   = "AUTH_AUDIENCE"
	claimsPathKey = "auth.jwt.claimsPath"
	claimsPathEnv = "AUTH_CLAIMS_PATH"
)

const (
	domainAuth0Key       = "auth.auth0.domain"
	domainAuth0Env       = "AUTH0_DOMAIN"
	clientIdAuth0Key     = "auth.auth0.clientId"
	clientIdAuth0Env     = "AUTH0_CLIENT_ID"
	clientSecretAuth0Key = "auth.auth0.clientSecret"
	clientSecretAuth0Env = "AUTH0_CLIENT_SECRET"
)

const (
	serverPort = "server.port"
)

type OIDCConfig struct {
	Issuer     string
	CacheTTL   time.Duration
	Audience   []string
	ClaimsPath string
}

type Auth0Config struct {
	Domain       string
	ClientId     string
	ClientSecret string
}

type DBConfig struct {
	Domain       string
	ClientId     string
	ClientSecret string
}

type AppConfig struct {
	auth0      *Auth0Config
	oidcConfig *OIDCConfig
	port       int
}

func init() {
	_ = viper.BindEnv(issuerKey, issuerEnv)
	_ = viper.BindEnv(audienceKey, audienceEnv)
	_ = viper.BindEnv(claimsPathKey, claimsPathEnv)

	_ = viper.BindEnv(domainAuth0Key, domainAuth0Env)
	_ = viper.BindEnv(clientIdAuth0Key, clientIdAuth0Env)
	_ = viper.BindEnv(clientSecretAuth0Key, clientSecretAuth0Env)
}

func NewConfig() *AppConfig {
	issuer := viper.GetString(issuerKey)
	audience := viper.GetStringSlice(audienceKey)
	claimsPath := viper.GetString(claimsPathKey)
	domainAuth0 := viper.GetString(domainAuth0Key)
	clientIdAuth0 := viper.GetString(clientIdAuth0Key)
	clientSecretAuth0 := viper.GetString(clientSecretAuth0Key)
	port := viper.GetInt(serverPort)

	if issuer == "" || audience == nil || claimsPath == "" || domainAuth0 == "" || clientIdAuth0 == "" || clientSecretAuth0 == "" {
		panic("cannot load configuration")
		return nil
	}

	oidcConfig := &OIDCConfig{
		Issuer:   issuer,
		Audience: audience,
		CacheTTL: 5 * time.Minute,
	}
	auth0Config := &Auth0Config{
		Domain:       domainAuth0,
		ClientId:     clientIdAuth0,
		ClientSecret: clientSecretAuth0,
	}

	if claimsPath != "" {
		oidcConfig.ClaimsPath = claimsPath
	}

	return &AppConfig{
		auth0:      auth0Config,
		oidcConfig: oidcConfig,
		port:       port,
	}
}

func SetAppConfig(config *AppConfig) {
	appConfig = config
}

func GetAppConfig() *AppConfig {
	return appConfig
}

func (a *AppConfig) GetServerPort() int {
	return a.port
}

func GetOIDCConfig() *OIDCConfig {
	return appConfig.oidcConfig
}

func GetAuth0Config() *Auth0Config {
	return appConfig.auth0
}
