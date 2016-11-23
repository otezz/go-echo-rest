package main

import (
    "fmt"
    "github.com/otezz/go-echo-rest/config"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/otezz/go-echo-rest/app/controllers"
)

var appConfig = config.Config.App
var jwtConfig = config.Config.JWT

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    //CORS
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
    }))

    // Routes
    e.GET("/", controllers.GetHome)

    v1 := e.Group("/v1")
    v1.POST("/login", controllers.PostLogin)

    v1.Use(middleware.JWT([]byte(jwtConfig.Secret)))

    // Users
    v1.GET("/users", controllers.AllUsers)
    v1.POST("/users", controllers.CreateUser)
    v1.GET("/users/:id", controllers.ShowUser)
    v1.PUT("/users/:id", controllers.UpdateUser)
    v1.DELETE("/users/:id", controllers.DeleteUser)

    // Articles
    v1.GET("/articles", controllers.AllArticles)
    v1.POST("/articles", controllers.CreateArticle)
    v1.GET("/articles/:id", controllers.ShowArticle)
    v1.PUT("/articles/:id", controllers.UpdateArticle)
    v1.DELETE("/articles/:id", controllers.DeleteArticle)


    // Server
    if err := e.Start(fmt.Sprintf("%s:%s", appConfig.HttpAddr, appConfig.HttpPort)); err != nil {
        e.Logger.Fatal(err.Error())
    }
}
