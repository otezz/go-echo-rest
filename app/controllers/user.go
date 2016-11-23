package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/otezz/go-echo-rest/app/models"
	"golang.org/x/crypto/bcrypt"
)

// get all users
func AllUsers(c echo.Context) error {
	var (
		users []models.User
		err   error
	)
	users, err = models.GetUsers()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, users)
}

// get one user
func ShowUser(c echo.Context) error {
	var (
		user models.User
		err  error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err = models.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, user)
}

//create user
func CreateUser(c echo.Context) error {
	user := new(models.User)

	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Password = hashPassword(c.FormValue("password"))

	err := models.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, user)
}

//update user
func UpdateUser(c echo.Context) error {
	// Parse the content
	user := new(models.User)

	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Password = hashPassword(c.FormValue("password"))

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update user data
	err = m.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

//delete user
func DeleteUser(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteUser()
	return c.JSON(http.StatusNoContent, err)
}

func hashPassword(input string) string {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
