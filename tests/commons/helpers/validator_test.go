package helpers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"github.com/fajardm/example_blog_api/commons/helpers"
)

type examplePost struct {
	Tile        string `validate:"required,gte=10"`
	Description string `validate:"required"`
}

func TestInitValidator(t *testing.T) {
	validator := helpers.Validator{}

	// ensure validator has function
	assert.True(t, reflect.TypeOf(validator.InitValidator).Kind() == reflect.Func, "InitValidator should a 'function'")
	assert.True(t, reflect.TypeOf(validator.Validate).Kind() == reflect.Func, "Validate should a 'function'")

	// create instance examplePost for test
	post := &examplePost{}

	// initialization validator
	validator.InitValidator()
	// validate examplePost
	validationErrors := validator.Validate(post)

	// ensure validation error not nil
	assert.NotNil(t, validationErrors, "validationErrors should not nil")
	assert.NotNil(t, validationErrors["Title"], "Title validation message should not nil")
	assert.NotNil(t, validationErrors["Description"], "Description validation message should not nil")
}

func TestInstanceValidator(t *testing.T) {
	assert.True(t, reflect.TypeOf(helpers.InstanceValidator).Kind() == reflect.Func, "GetInstanceValidator should a 'function'")

	validator := helpers.InstanceValidator()

	// ensure validator is singleton
	assert.Equal(t, validator, helpers.InstanceValidator(), "validator should singleton")

	newValidator := helpers.Validator{}
	newValidator.InitValidator()
	assert.NotEqual(t, newValidator, validator, "validator should not equal newValidator")

	// create instance examplePost for test
	post := &examplePost{Tile: "post"}
	validationErrors := validator.Validate(post)

	// ensure validation error not nil
	assert.NotNil(t, validationErrors, "validationErrors should not nil")
	assert.NotNil(t, validationErrors["Title"], "Title validation message should not nil")
	assert.NotNil(t, validationErrors["Description"], "Description validation message should not nil")
}
