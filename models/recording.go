package models

import (
    _ "errors"
    _ "fmt"
    _ "strconv"
    _ "time"
    "github.com/bigbluebutton/bbb-api-recordings/utils"
)

var (
    RecordingList map[string]*Recording
)

func init() {
    RecordingList = make(map[string]*Recording)
    r1 := Recording{"recording_1", "A Nice Recording", "room_1", true}
    RecordingList["recording_1"] = &r1
    r2 := Recording{"recording_2", "Another Nice Recording", "room_1", false}
    RecordingList["recording_2"] = &r2
    r3 := Recording{"recording_3", "Yet Another", "room_2", true}
    RecordingList["recording_3"] = &r3
}

type Recording struct {
    MeetingId string `json:"meetingId"`
    Name      string `json:"name"`
    RoomId    string `json:"roomId"`
    Published bool   `json:"published"`
}

func GetAllRecordings(filters *RecordingFilters) []*Recording {
    ret := []*Recording{}

    // no filters, return all
    if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
        for _, rec := range RecordingList {
            ret = append(ret, rec)
        }
        return ret
    }

    // at least one filter selected
    for _, rec := range RecordingList {
        roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
        meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
            ret = append(ret, rec)
        }
    }

    return ret
}

func DeleteAllRecordings(filters *RecordingFilters) ([]*Recording, []APIError) {
    errors := []APIError{}
    recs := []*Recording{}

    // no filters, be safe and don't do anything
    if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
        return recs, errors
    }

    // at least one filter selected
    for _, rec := range GetAllRecordings(nil) {
        roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
        meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
            if rec, err := DeleteRecording(rec.MeetingId); err != nil {
                errors = append(errors, *err)
            } else {
                recs = append(recs, rec)
            }
        }
    }

    return recs, errors
}

func UpdateAllRecordings(filters *RecordingFilters, params *Recording) ([]*Recording, []APIError) {
    errors := []APIError{}
    recs := []*Recording{}

    // no filters, be safe and don't do anything
    if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
        return recs, errors
    }

    // at least one filter selected
    for _, rec := range GetAllRecordings(nil) {
        roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
        meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
            if rec, err := UpdateRecording(rec.MeetingId, params); err != nil {
                errors = append(errors, *err)
            } else {
                recs = append(recs, rec)
            }
        }
    }

    return recs, errors
}

func DeleteRecording(uid string) (rec *Recording, err *APIError) {
    rec = RecordingList[uid]
    delete(RecordingList, uid)
    return rec, nil
}

func UpdateRecording(uid string, params *Recording) (a *Recording, err *APIError) {
    if r, ok := RecordingList[uid]; ok {
        if params.Name != "" {
            r.Name = params.Name
        }
        // if params.Published != nil {
        //  r.Published = params.Published
        // }
        return r, nil
    }
    return nil, &APIError{"NotFound", "Recording does not exist", &uid}
}
