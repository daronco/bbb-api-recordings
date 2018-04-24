package controllers

import (
	"../models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// Operations about Recordings
type RecordingController struct {
	beego.Controller
}

// @Title Get
// @Description get all Recordings
// @Success 200 {object} models.Recording
// @router / [get]
func (c *RecordingController) Get() {
	recordings := models.GetAllRecordings()
	c.Data["json"] = recordings
	c.ServeJson()
}

// @Title Get
// @Description get all Recordings
// @Success 200 {object} models.Recording
// @router / [get]
func (c *RecordingController) Delete() {
	result, err := models.DeleteAllRecordings()
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = result
	}
	c.ServeJson()
}

func (c *RecordingController) Update() {
	var recording models.Recording
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &recording)
	if err != nil {
		fmt.Println("Error parsing", err)
	}
	result, err2 := models.UpdateAllRecordings(&recording)
	if err2 != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = result
	}
	c.ServeJson()
}
