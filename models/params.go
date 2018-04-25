package models

import (
    _ "encoding/json"
)

// Parameters accepted when receiving requests
type RecordingParams struct {
    Filters RecordingFilters `json:"filters"`
    Attributes Recording     `json:"attributes"`
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
