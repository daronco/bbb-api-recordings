package controllers

import (
    "github.com/bigbluebutton/bbb-api-recordings/models"

    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error404() {
    err := models.APIError{"NotFound", "API method not found", nil}
    errs := []models.APIError{err}
    response := models.APIErrorResponse{errs}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *ErrorController) Error500() {
    err := models.APIError{"ServerError", "Internal server error", nil}
    errs := []models.APIError{err}
    response := models.APIErrorResponse{errs}

    c.Data["json"] = response
    c.ServeJSON()
}
