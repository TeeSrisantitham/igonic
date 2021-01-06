package config

import (
	"fmt"
	"testing"
	"os"

	"github.com/openware/igonic/helper"
	"github.com/openware/igonic/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func comparePage(p1 *models.Page, p2 *models.Page) bool {
	if p1.Path == p2.Path && 
	p1.Lang == p2.Lang &&
	p1.Title == p2.Title &&
	p1.Keywords == p2.Keywords &&
	p1.Description == p2.Description &&
	p1.Body == p2.Body {
		return true
	}
	return false
}

func compareArticle(a1 *models.Article, a2 *models.Article) bool {
	if a1.Slug == a2.Slug &&
	a1.AuthorUID == a2.AuthorUID && 
	a1.Lang == a2.Lang &&
	a1.Title == a2.Title &&
	a1.Keywords == a2.Keywords &&
	a1.Description == a2.Description &&
	a1.Body == a2.Body {
		return true
	}
	return false
}

func TestSuccess(t *testing.T) {
	// connect to test db
	ctx := helper.GetTestContext()

	// change root path to project root
	os.Chdir("..")
	
	migrateDB(ctx)

	// prepare expect result
	pageExpected := models.Page{
		Path: "/terms",
		Lang: "EN",
		Title:"Term of services",
		Keywords:"terms, tos",
		Description: "Term of services",
		Body: fmt.Sprintf("# Term of services\nThis is an example of page\n"),
	} 

	articleExpected := models.Article{
		Slug: "hello",
		AuthorUID: "ABC00001",
		Lang: "EN",
		Title:"Welcome to openware gin skel",
		Keywords:"Gin, Gonic, Framework",
		Description: "this is an example of article using the openware skeleton framework",
		Body: fmt.Sprintf("# Welcome the openware Gin skeleton\nThis is an example of article\n\n## What syntax to use to write an article?\nArticles are written in [Markdown](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet).\nThis is a very simple language which allows to format the content of an article, add links, images..\n"),
	} 

	// call function seeds
	env := "test.yml"
	err := Seeds(ctx, env)
	require.Equal(t, nil, err, "It should no error")

	// run test
	t.Run("Page", func(t *testing.T) {
		var page models.Page
		result := ctx.First(&page)
		assert.EqualValues(t, 1, result.RowsAffected)
		assert.EqualValues(t, true, comparePage(&pageExpected, &page))
	})

	t.Run("Article", func(t *testing.T) {
		var article models.Article
		result := ctx.First(&article)
		assert.EqualValues(t, 1, result.RowsAffected)
		assert.EqualValues(t, true, compareArticle(&articleExpected, &article))
	})
}

func TestFail_FileNotFound(t *testing.T) {
	// connect to test db
	ctx := helper.GetTestContext()
	
	migrateDB(ctx)
	
	// call function
	env := "empty_test.yml"
	err := Seeds(ctx, env)

	// run test
	assert.Error(t, err)
}

func TestFail_DB(t *testing.T) {
	// connect to test db
	ctx := helper.GetTestContext()
	
	// call function
	env := "test.yml"
	err := Seeds(ctx, env)

	// run test
	assert.Error(t, err)
}

func migrateDB(ctx *gorm.DB) {
	ctx.AutoMigrate(&models.Article{})
	ctx.AutoMigrate(&models.Page{})
}
