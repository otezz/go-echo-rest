package controllers

import (
    "net/http"

    "github.com/labstack/echo"
)

func GetHome(c echo.Context) error {
    return c.String(http.StatusOK, "oke")
}
