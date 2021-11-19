package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CheckStruct(src interface{}) error {
	err := validate.Struct(src)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Errorf("invalid struct: %w", err)
	}

	errs := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		// fmt.Println(err.Namespace())
		// fmt.Println(err.Field())
		// fmt.Println(err.StructNamespace())
		// fmt.Println(err.StructField())
		// fmt.Println(err.Tag())
		// fmt.Println(err.ActualTag())
		// fmt.Println(err.Kind())
		// fmt.Println(err.Type())
		// fmt.Println(err.Value())
		// fmt.Println(err.Param())
		errs = append(errs, err.Error())
	}

	return fmt.Errorf("invalid struct: %s", strings.Join(errs, " "))
}

func CheckValue(src interface{}, tag string) error {
	err := validate.Var(src, tag)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Errorf("invalid value: %w", err)
	}

	errs := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, err.Error())
	}

	return fmt.Errorf("invalid value: %s", strings.Join(errs, " "))
}
