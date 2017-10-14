package controllers

import {
	"github.com/labstack/echo"
	"net/http"
}

// Create User
func CreateUser(c echo.Context) error {
	user := new(models.User)
	var err error
	return c.JSON(http.StatusCreated, user)
}