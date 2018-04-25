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

	// parse accepted URL parameters
	meetingIds := c.GetStrings("meetingId")
	roomIds := c.GetStrings("roomId")

	// parse request body
	var params models.RecordingIndexParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// fmt.Println("Parsed", params)
	if err != nil {
		fmt.Println("Error parsing request body", err)
	}

	// give priority to parameters set in the URL
	if len(meetingIds) > 0 {
		params.Filters.MeetingIds = meetingIds
	}
	if len(roomIds) > 0 {
		params.Filters.RoomIds = roomIds
	}

	recordings := models.GetAllRecordings(&params.Filters)
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
	// var anyJson map[string]interface{}
	// err := json.Unmarshal(c.Ctx.Input.RequestBody, &anyJson)
	// fmt.Println("Parsed", anyJson)

	// meetingIds := make([]string, 0)
	// var meetingIds string
	// c.Ctx.Input.Bind(&meetingIds, "meetingId")

	// var body models.RecordingRequestBody
	// err := c.ParseForm(&body)
	// fmt.Println("Parsed", body)

	// parse accepted URL parameters
	meetingIds := c.GetStrings("meetingId")
	roomIds := c.GetStrings("roomId")

	// parse request body
	var params models.RecordingUpdateParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	fmt.Println("Parsed", params)
	if err != nil {
		fmt.Println("Error parsing request body", err)
	}

	// give priority to parameters set in the URL
	if len(meetingIds) > 0 {
		params.Filters.MeetingIds = meetingIds
	}
	if len(roomIds) > 0 {
		params.Filters.RoomIds = roomIds
	}

	c.Data["json"] = params

	// result, err2 := models.UpdateAllRecordings(&recording)
	// if err2 != nil {
	// 	c.Data["json"] = err
	// } else {
	// 	c.Data["json"] = result
	// }
	c.ServeJson()
}
