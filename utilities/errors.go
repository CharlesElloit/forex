package utilities

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
)

func CreateError(statusCode int, title string, detail string, ctx iris.Context) {
	ctx.StopWithProblem(statusCode, iris.NewProblem().Title(title).Detail(detail))
}

func CreateInternalServerError(ctx iris.Context) {
	CreateError(
		iris.StatusInternalServerError,
		"Internal Server Error",
		"Internal Server Error",
		ctx,
	)
}

func HandleValidationError(err error, ctx iris.Context) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		validationErrors := WrapValidationErrors(errors)
		// fmt.Println("ValidationErrors: ", validationErrors)
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().
				Title("Validation errors").
				Detail("One or more fields failed to be validated, please check and try again.").
				Key("errors", validationErrors),
		)
		return
	}
	CreateInternalServerError(ctx)
}

func WrapValidationErrors(errors validator.ValidationErrors) []validationError {
	validationErrors := make([]validationError, 0, len(errors))
	for _, validationErr := range errors {
		validationErrors = append(validationErrors, validationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}
	return validationErrors
}

type validationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}
