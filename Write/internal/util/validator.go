package util

import (
	"github/SLANGERES/CQRS/Write/internal/models"

	"github.com/go-playground/validator/v10"
)

// Initialize a single validator instance
var validate = validator.New()

// ValidateReqBody validates the Blog struct
func ValidateReqBody(blog models.Blog) error {
	return validate.Struct(blog)
}
