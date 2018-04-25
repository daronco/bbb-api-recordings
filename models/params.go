package models

import (
	_ "encoding/json"
)

// Parameters accepted when receiving a request to update a recording
type RecordingUpdateParams struct {
	Filters RecordingFilters `json:"filters"`
	Attributes Recording     `json:"attributes"`
}

// Parameters accepted when receiving a request to list recordings
type RecordingIndexParams struct {
	Filters RecordingFilters `json:"filters"`
}

// Parameters accepted when filtering recordings (used in multiple routes)
type RecordingFilters struct {
	MeetingIds []string `json:"meetingId"`
	RoomIds []string    `json:"roomId"`
}
