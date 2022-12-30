package helpers

import (
	"anime-hentai-backend/models"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type ValidateStructs interface {
	models.User | interface{}
}

var validate = validator.New()

func ValidateStruct[T ValidateStructs](myStruct T) []*ErrorResponse {
	var errors []*ErrorResponse
	if err := validate.Struct(myStruct); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
