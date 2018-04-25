package models

import (
	_ "encoding/json"
)

// Parameters accepted when receiving a request to list recordings
type RecordingIndexParams struct {
	Filters RecordingFilters `json:"filters"`
}

// Parameters accepted when receiving a request to update a recording
type RecordingUpdateParams struct {
	Filters RecordingFilters `json:"filters"`
	Attributes Recording     `json:"attributes"`
}

// Parameters accepted when receiving a request to delete recordings
type RecordingDeleteParams struct {
	Filters RecordingFilters `json:"filters"`
}

// Default response
type RecordingResponse struct {
	Data map[string]*Recording `json:"data"`
	Errors map[string]*error   `json:"errors"`
}

// Parameters accepted when filtering recordings (used in multiple routes)
type RecordingFilters struct {
	MeetingIds []string `json:"meetingId"`
	RoomIds []string    `json:"roomId"`
}
