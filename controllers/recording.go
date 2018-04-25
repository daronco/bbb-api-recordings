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

func (c *RecordingController) ParseParams() *models.RecordingParams {
    var params models.RecordingParams

    // parse accepted URL parameters
    meetingIds := c.GetStrings("meetingId")
    roomIds := c.GetStrings("roomId")

    // parse request body
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

    return &params
}

// @Title Get Recordings
// @Description get a list of Recordings
// @Success 200 {object} models.Recording
// @router / [get]
func (c *RecordingController) Get() {
    params := c.ParseParams()
    recs := models.GetAllRecordings(&params.Filters)
    if len(recs) == 0 { recs = nil }

    response := models.RecordingResponse{recs, nil}
    c.Data["json"] = response
    c.ServeJson()
}

// @Title Delete Recordings
// @Description delete one or more Recordings
// @Success 200 {object} models.Recording
// @router / [delete]
func (c *RecordingController) Delete() {
    params := c.ParseParams()
    filters := &params.Filters
    var response models.RecordingResponse

    if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
        err := models.APIError{"noFilters", "Request aborted because no filters were provided", nil}
        errs := []models.APIError{err}
        response = models.RecordingResponse{nil, errs}
    } else {
        recs, errs := models.DeleteAllRecordings(filters)
        if len(recs) == 0 { recs = nil }
        if len(errs) == 0 { errs = nil }

        response = models.RecordingResponse{recs, errs}
    }
    c.Data["json"] = response
    c.ServeJson()
}

// @Title Update Recordings
// @Description update one or more Recordings
// @Success 200 {object} models.Recording
// @router / [patch]
func (c *RecordingController) Update() {
    params := c.ParseParams()
    filters := &params.Filters
    var response models.RecordingResponse

    if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
        err := models.APIError{"noFilters", "Request aborted because no filters were provided", nil}
        errs := []models.APIError{err}
        response = models.RecordingResponse{nil, errs}
    } else {
        recs, errs := models.UpdateAllRecordings(&params.Filters, &params.Attributes)
        if len(recs) == 0 { recs = nil }
        if len(errs) == 0 { errs = nil }

        response = models.RecordingResponse{recs, errs}
    }
    c.Data["json"] = response
    c.ServeJson()
}
