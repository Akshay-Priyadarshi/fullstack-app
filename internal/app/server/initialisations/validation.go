package initialisations

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func InitValidation(translator *ut.Translator) *validator.Validate {
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, *translator)
	return validate
}
