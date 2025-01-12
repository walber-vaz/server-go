package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR2 "github.com/go-playground/validator/v10/translations/pt_BR"
	"server-go/internal/config/rest_errors"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		br := pt_BR.New()
		un := ut.New(br, br)
		transl, _ = un.GetTranslator("pt_BR")
		err := pt_BR2.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validationErr error) *rest_errors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_errors.NewBadRequestError("Ops! O campo " + jsonErr.Field + " está com o tipo de dado incorreto.")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsValidations []rest_errors.Err

		for _, err := range validationErr.(validator.ValidationErrors) {
			err := rest_errors.Err{
				Field:   err.Field(),
				Message: err.Translate(transl),
			}

			errorsValidations = append(errorsValidations, err)
		}

		return rest_errors.NewBadRequestErrorValidation("Ops! Ocorreu um erro de validação.", errorsValidations)
	} else {
		return rest_errors.NewBadRequestError("Ops! Ocorreu um erro de validação.")
	}
}
