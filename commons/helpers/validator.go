package helpers

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/locales/en"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	"log"
	"sync"
)

type Validator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

func (v *Validator) InitValidator() *Validator {
	v.validate = validator.New()
	v.setTrans()
	return v
}

// we don't need setTrans for public access
func (v *Validator) setTrans() {
	// ensure validate is not nil to working with translator
	if v.validate != nil {
		newEn := en.New()
		v.uni = ut.New(newEn, newEn)
		v.trans, _ = v.uni.GetTranslator("en")
		enTranslations.RegisterDefaultTranslations(v.validate, v.trans)
	} else {
		log.Panic("validate is nil")
	}
}

// this function to validate struct
func (v *Validator) Validate(data interface{}) map[string]string {
	err := v.validate.Struct(data)
	validationErrors := err.(validator.ValidationErrors)

	messages := make(map[string]string)

	for _, e := range validationErrors {
		messages[e.Field()] = e.Translate(v.trans)
	}

	return messages
}

var (
	instance *Validator
	once     sync.Once
)

// access validator singleton
func InstanceValidator() *Validator {
	once.Do(func() {
		instance = &Validator{}
		instance.InitValidator()
	})

	return instance
}
