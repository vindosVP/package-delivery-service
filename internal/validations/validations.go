package validations

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidations() error {
	enval := en.New()
	uni = ut.New(enval, enval)

	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	err := enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	return nil
}

func UniversalValidation(body interface{}) (bool, []*ValidationErrorResponse) {
	var errors []*ValidationErrorResponse
	if err := validate.Struct(body); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors = append(errors, &ValidationErrorResponse{
				Field:   e.Field(),
				Message: e.Translate(trans),
			})
		}

		return false, errors
	}
	return true, nil
}
