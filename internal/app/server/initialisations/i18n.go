package initialisations

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

func InitI18n() *ut.Translator {
	en := en.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")
	return &translator
}
