package validators

import (
	"github.com/go-playground/validator/v10"
	"os"
	"reflect"
)

func AddYamlFileValidator(validate *validator.Validate) error {
	return validate.RegisterValidation("yaml_file", yamlFileValidator)
}

func yamlFileValidator(fl validator.FieldLevel) bool {
	fieldVal := fl.Field()
	if fieldVal.Kind() == reflect.Pointer {
		fieldVal = fieldVal.Elem()
	}

	if fieldVal.Kind() != reflect.String {
		return false
	}

	val := fieldVal.String()

	f, err := os.Open(val)
	if err != nil {
		return false
	}

	defer func() {
		_ = f.Close()
	}()

	return isYaml(f)
}
