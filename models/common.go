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

// A generic API error response
type APIErrorResponse struct {
    Errors []APIError `json:"errors"`
}

// Default response for recording routes
type RecordingResponse struct {
    Data []*Recording `json:"data"`
    Errors []APIError `json:"errors"`
}

// Parameters accepted when filtering recordings (used in multiple routes)
type RecordingFilters struct {
    MeetingIds []string `json:"meetingId"`
    RoomIds []string    `json:"roomId"`
}

// Parameters accepted when receiving requests for recordings
type RecordingParams struct {
    Filters RecordingFilters `json:"filters"`
    Attributes Recording     `json:"attributes"`
}
