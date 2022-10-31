package auth

import (
	"github.com/labstack/echo/v4"
	"waas-service/dto"
)

type Context struct {
	echo.Context
	Claims *dto.UserClaims
}
