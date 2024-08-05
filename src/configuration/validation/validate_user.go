package validation

import (
	"encoding/json"
	"errors"
	"log"

	"example.com/mod/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validation = validator.New()
	transl     ut.Translator
)

func init() {

	log.Printf("Init validation")
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	log.Printf("Error ValidateUserError: %+v", validation_err)

	if ok := errors.As(validation_err, &jsonErr); ok {

		log.Printf("Error trying to convert field: %+v", jsonErr)
		return rest_err.NewBadRequestError("Invalid field type")

	} else if ok := errors.As(validation_err, &jsonValidationError); ok {

		log.Printf("Error trying to convert field: %+v", jsonValidationError)
		errorsCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			causes := rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}
			log.Printf("Error trying to convert field cause: %+v", causes)
			errorsCauses = append(errorsCauses, causes)
		}

		return rest_err.NewBadRequestValidationError("Some field are invalid", errorsCauses)

	} else {

		log.Printf("Error trying to convert field: %+v", validation_err)
		return rest_err.NewBadRequestError("Error trying to convert field")
	}
}
