package middleware

import (
	"github.com/clubo-app/packages/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ParseAccessTokenMapClaims(claims jwt.MapClaims) types.AccessTokenPayload {
	payload := types.AccessTokenPayload{}

	if iss, ok := claims["iss"].(string); ok {
		payload.Iss = iss
	}

	if sub, ok := claims["sub"].(string); ok {
		payload.Sub = sub
	}

	if iat, ok := claims["iat"].(float64); ok {
		payload.Iat = iat
	}

	if exp, ok := claims["exp"].(int64); ok {
		payload.Exp = exp
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

func ParseUser(c *fiber.Ctx) types.AccessTokenPayload {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return types.AccessTokenPayload{}
	}
	claims := user.Claims.(jwt.MapClaims)

	payload := ParseAccessTokenMapClaims(claims)

	return payload
}
