package middleware

import (
	"github.com/clubo-app/packages/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ParseMapClaims(claims jwt.MapClaims) types.JwtPayload {
	payload := types.JwtPayload{}

	if iss, ok := claims["iss"].(string); ok {
		payload.Iss = iss
	}

	if sub, ok := claims["sub"].(string); ok {
		payload.Sub = sub
	}

	if iat, ok := claims["iat"].(float64); ok {
		payload.Iat = iat
	}

	if pStr, ok := claims["provider"].(string); ok {
		payload.Provider = types.ProviderFromString(pStr)
	}

	if rStr, ok := claims["role"].(string); ok {
		payload.Role = types.RoleFromString(rStr)
	}

	if eVerified, ok := claims["emailVerified"].(bool); ok {
		payload.EmailVerified = eVerified
	}

	return payload
}

func ParseUser(c *fiber.Ctx) types.JwtPayload {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return types.JwtPayload{}
	}
	claims := user.Claims.(jwt.MapClaims)

	payload := ParseMapClaims(claims)

	return payload
}
