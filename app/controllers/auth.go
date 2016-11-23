package controllers

import (
    "github.com/labstack/echo"
    //"github.com/dgrijalva/jwt-go"
    //"time"
    "net/http"
    //"os/user"
    "github.com/otezz/go-echo-rest/app/models"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "github.com/otezz/go-echo-rest/config"
)

var jwtConfig = config.Config.JWT

func PostLogin(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    user, err := models.GetUserByUsername(username)
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, err)
    }

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "message": "Invalid username / password",
        })
    }

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["username"] = user.Username
    claims["email"] = user.Email
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token and send it as response.
    t, err := token.SignedString([]byte(jwtConfig.Secret))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}
