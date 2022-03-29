package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type formRequestErrors struct {
	errorMaps map[string]string
}

func (f formRequestErrors) Validate(model interface{}, err error) map[string]string {
	errs := map[string]string{}
	fields := map[string]FormRequestResult{}

	if _, ok := err.(validator.ValidationErrors); ok {
		types := reflect.TypeOf(model)

		for i := 0; i < types.NumField(); i++ {
			field := types.Field(i)

			jsonTag := map[bool]string{
				true:  field.Name,
				false: field.Tag.Get("json"),
			}[field.Tag.Get("json") == ""]

			messageTag := field.Tag.Get("msg")
			msg := f.getErrorMessage(messageTag)

			fields[field.Name] = FormRequestResult{
				Field:   field.Name,
				JsonTag: jsonTag,
				Message: msg,
			}
		}

		for _, e := range err.(validator.ValidationErrors) {
			if field, ok := fields[e.Field()]; ok {
				errs[field.JsonTag] = map[bool]string{
					true:  field.Message,
					false: e.Error(),
				}[field.Message != ""]
			}
		}
	}

	return errs
}

func (f formRequestErrors) getErrorMessage(key string) string {
	if value, ok := f.errorMaps[key]; ok {
		return value
	}

	return key
}

func NewFormRequest(errors map[string]string) FormRequest {
	return &formRequestErrors{errorMaps: errors}
}
