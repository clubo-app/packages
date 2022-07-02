package types

type MessageRes struct {
	Message string `json:"message" validate:"required"`
}

type AccessTokenPayload struct {
	Iss           string   `json:"iss"`
	Sub           string   `json:"sub"`
	Iat           float64  `json:"iat"`
	Exp           int64    `json:"exp"`
	Role          Role     `json:"role"`
	EmailVerified bool     `json:"emailVerified"`
	Provider      Provider `json:"provider"`
}
