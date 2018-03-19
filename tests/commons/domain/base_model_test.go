package domain

import (
	"github.com/jinzhu/gorm"
	"testing"
	"github.com/fajardm/example_blog_api/commons/helpers"
	"os"
	"github.com/fajardm/example_blog_api/commons/domain"
	"github.com/stretchr/testify/assert"
)

// this class just for testing
type ExamplePost struct {
	domain.BaseModel
	title string
}

var db *gorm.DB

// run before main testing
func Setup() {
	// create database connection than migrate table
	db, _ = helpers.InitConnection(os.Getenv("DB_DIALECT"), os.Getenv("DB_URL"))
	db.CreateTable(ExamplePost{})
}

// run after main testing
func Teardown() {
	// drop table ExamplePost than close database
	db.DropTable(ExamplePost{})
	db.Close()
}

func TestMain(m *testing.M) {
	Setup()
	retCode := m.Run()
	Teardown()
	os.Exit(retCode)
}

func TestCreateRecord(t *testing.T) {
	// create object example post
	examplePost := ExamplePost{
		title: "POST 1",
	}

	// store object examplePost to database
	db.Create(&examplePost)

	// ensure examplePost not nil
	assert.NotNil(t, examplePost, "examplePost should not nil")
	assert.NotNil(t, examplePost.ID, "examplePost.ID should not nil")
	assert.NotNil(t, examplePost.CreatedAt, "examplePost.CreatedAt should not nil")
	assert.NotNil(t, examplePost.UpdatedAt, "examplePost.UpdatedAt should not nil")
	assert.Nil(t, examplePost.DeletedAt, "examplePost.DeletedAt should nil")
}
