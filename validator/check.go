package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Check(src interface{}) error {
	err := validate.Struct(src)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Errorf("Invalid: %w", err)
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

	return fmt.Errorf("Invalid: %s", strings.Join(errs, " "))
}
