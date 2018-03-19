package helpers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"github.com/fajardm/example_blog_api/commons/helpers"
	"os"
)

func TestInitConnection(t *testing.T) {
	// ensure InitConnection is a 'function'
	assert.True(t, reflect.TypeOf(helpers.InitConnection).Kind() == reflect.Func, "InitConnection should a 'function'")

	// create database connection than defer db close
	db, err := helpers.InitConnection(os.Getenv("DB_DIALECT"), os.Getenv("DB_URL"))
	defer db.Close()

	// ensure connection working properly
	assert.Nil(t, err, "Connection should not error")
	assert.NotNil(t, db, "Connection should working")
}

func TestGetConnection(t *testing.T) {
	// ensure GetConnection is a 'function'
	assert.True(t, reflect.TypeOf(helpers.GetConnection).Kind() == reflect.Func, "GetConnection should a 'function'")

	// create database connection
	helpers.InitConnection(os.Getenv("DB_DIALECT"), os.Getenv("DB_URL"))

	// get database connection than defer db close
	db := helpers.GetConnection()
	defer db.Close()

	// ensure connection working properly
	assert.NotNil(t, db, "Connection should working")
}
