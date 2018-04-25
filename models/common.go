package models

import (
    _ "encoding/json"
)

// Description of an error
type APIError struct {
    Code string        `json:"code"`
    Description string `json:"description"`
    Source *string     `json:"source"`
}

// Parameters accepted when receiving requests
type RecordingParams struct {
    Filters RecordingFilters `json:"filters"`
    Attributes Recording     `json:"attributes"`
}

// Default response
type RecordingResponse struct {
    Data map[string]*Recording `json:"data"`
    Errors []APIError          `json:"errors"`
}

// Parameters accepted when filtering recordings (used in multiple routes)
type RecordingFilters struct {
    MeetingIds []string `json:"meetingId"`
    RoomIds []string    `json:"roomId"`
}
