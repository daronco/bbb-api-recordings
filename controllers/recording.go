package controllers

import (
	"../models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type RecordingController struct {
	beego.Controller
}

// @Title Get Recordings
// @Description get a list of Recordings
// @Success 200 {object} models.Recording
// @router / [get]
func (c *RecordingController) Get() {

	// parse accepted URL parameters
	meetingIds := c.GetStrings("meetingId")
	roomIds := c.GetStrings("roomId")

	// parse request body
	var params models.RecordingIndexParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
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

// @Title Delete Recordings
// @Description delete one or more Recordings
// @Success 200 {object} models.Recording
// @router / [delete]
func (c *RecordingController) Delete() {

	// parse accepted URL parameters
	meetingIds := c.GetStrings("meetingId")
	roomIds := c.GetStrings("roomId")

	// parse request body
	var params models.RecordingDeleteParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
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

	// TODO: if there are no filters, return an error
	recs, errs := models.DeleteAllRecordings(&params.Filters)

	response := models.RecordingResponse{recs, errs}
	c.Data["json"] = response
	c.ServeJson()
}

// @Title Update Recordings
// @Description update one or more Recordings
// @Success 200 {object} models.Recording
// @router / [patch]
func (c *RecordingController) Update() {

	// parse accepted URL parameters
	meetingIds := c.GetStrings("meetingId")
	roomIds := c.GetStrings("roomId")

	// parse request body
	var params models.RecordingUpdateParams
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
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

	// TODO: if there are no filters, return an error
	recs, errs := models.UpdateAllRecordings(&params.Filters, &params.Attributes)

	response := models.RecordingResponse{recs, errs}
	c.Data["json"] = response
	c.ServeJson()
}
