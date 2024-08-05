package validation

import (
	"encoding/json"
	"errors"

	"example.com/mod/src/configuration/logger"
	"example.com/mod/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/zap"
)

var (
	Validation = validator.New()
	transl     ut.Translator
)

func init() {
	logger.Info("Init Validate user", zap.String("journey", "ValidateUser"))
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
		logger.Info("Validator engine started", zap.String("journey", "ValidateUser"))
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	logger.Error("Error trying to validate user info", validation_err, zap.String("journey", "ValidateUser"))

	if ok := errors.As(validation_err, &jsonErr); ok {

		logger.Error("Error trying to convert field", jsonErr, zap.String("journey", "ValidateUser"))
		return rest_err.NewBadRequestError("Invalid field type")

	} else if ok := errors.As(validation_err, &jsonValidationError); ok {

		logger.Error("Error trying to convert field", jsonValidationError, zap.String("journey", "ValidateUser"))
		errorsCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			causes := rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}

			logger.Error("Error trying to convert field cause"+causes.Field+causes.Message, nil, zap.String("journey", "ValidateUser"))
			errorsCauses = append(errorsCauses, causes)
		}

		return rest_err.NewBadRequestValidationError("Some field are invalid", errorsCauses)

	} else {

		logger.Error("Error trying to convert field", validation_err, zap.String("journey", "ValidateUser"))
		return rest_err.NewBadRequestError("Error trying to convert field")
	}
}
