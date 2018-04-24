package controllers

import (
	"../models"
	_ "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Recordings
type RecordingController struct {
	beego.Controller
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *RecordingController) GetAll() {
	recordings := models.GetAllRecordings()
	u.Data["json"] = recordings
	u.ServeJson()
}
