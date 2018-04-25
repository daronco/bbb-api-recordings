package controllers

import (
    "../models"

    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error404() {
    err := models.APIError{"NotFound", "Method not found", nil}
    errs := []models.APIError{err}
    response := models.APIErrorResponse{errs}

    c.Data["json"] = response
    c.ServeJson()
}

func (c *ErrorController) Error500() {
    err := models.APIError{"InternalServerError", "Internal server error", nil}
    errs := []models.APIError{err}
    response := models.APIErrorResponse{errs}

    c.Data["json"] = response
    c.ServeJson()
}
