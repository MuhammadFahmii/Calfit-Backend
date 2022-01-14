package middlewares

import (
	"CalFit/controllers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Admin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)
			if claims.Admin {
				return hf(c)
			} else {
				return controllers.ErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}

func Superadmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)
			if claims.Superadmin {
				return hf(c)
			} else {
				return controllers.ErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}
