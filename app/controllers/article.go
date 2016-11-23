package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/otezz/go-echo-rest/app/models"
)

// get all articles
func AllArticles(c echo.Context) error {
	var (
		articles []models.Article
		err      error
	)
	articles, err = models.GetArticles()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, articles)
}

// get one article
func ShowArticle(c echo.Context) error {
	var (
		article models.Article
		err     error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	article, err = models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, article)
}

//create article
func CreateArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	article := new(models.Article)

	article.UserID = userId
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	err := models.CreateArticle(article)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, article)
}

//update article
func UpdateArticle(c echo.Context) error {
	// Parse the content
	article := new(models.Article)

	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update article data
	err = m.UpdateArticle(article)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

//delete article
func DeleteArticle(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteArticle()
	return c.JSON(http.StatusNoContent, err)
}
