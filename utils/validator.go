package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type ApiError struct {
	Field string
	Msg   string
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorTag := ve[0].Tag()
			errorField := ve[0].Field()

			return echo.NewHTTPError(http.StatusBadRequest, msgForTag(errorTag, errorField))
		}
		// Optionally, you could return the error to give each route more control over the status code
		return nil
	}
	return nil
}

func msgForTag(tag string, field string) string {
	var b strings.Builder

	switch tag {
	case "required":
		switch field {
		case "PId":
			b.WriteString("Party Id")
		case "UId":
			b.WriteString("User Id")
		case "CId":
			b.WriteString("Comment Id")
		default:
			b.WriteString(field)
		}
		b.WriteString(" is required")
		return b.String()
	case "email":
		return "Invalid Email"
	case "gte":
		if field == "Password" {
			return "Password needs to be at least 8 characters long"
		}
		b.WriteString(field)
		b.WriteString(" needs to be longer")
		return b.String()
	}
	return "Invalid body format"
}
