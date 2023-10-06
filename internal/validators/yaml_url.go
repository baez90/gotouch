package validators

import (
	"context"
	"github.com/denizgursoy/gotouch/internal/auth"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

func AddYamlUrlValidator(validate *validator.Validate) error {
	return validate.RegisterValidationCtx("yaml_url", yamlUrlValidator)
}

func yamlUrlValidator(ctx context.Context, fl validator.FieldLevel) bool {
	fieldVal := fl.Field()
	if fieldVal.Kind() == reflect.Pointer {
		fieldVal = fieldVal.Elem()
	}

	if fieldVal.Kind() != reflect.String {
		return false
	}

	val := fieldVal.String()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, val, nil)
	if err != nil {
		return false
	}

	httpClient := auth.NewAuthenticatedHTTPClient()
	resp, err := httpClient.Do(req)
	if err != nil {
		return false
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return false
	}

	return isYaml(resp.Body)
}
