package types

type MessageRes struct {
	Message string `json:"message" validate:"required"`
}

type JwtPayload struct {
	Iss           string   `json:"iss"`
	Sub           string   `json:"sub"`
	Iat           float64  `json:"iat"`
	Role          Role     `json:"role"`
	EmailVerified bool     `json:"emailVerified"`
	Provider      Provider `json:"provider"`
}
