package models

import (
	"errors"
	"../utils"
	_ "fmt"
	_ "strconv"
	_ "time"
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

func GetAllRecordings(filters *RecordingFilters) map[string]*Recording {
	// no filters, return all
	if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
		return RecordingList
	}

	// at least one filter selected
	ret := make(map[string]*Recording)
    for id, rec := range RecordingList {
		roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
		meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
			ret[id] = rec
		}
    }
	return ret
}

func DeleteAllRecordings(filters *RecordingFilters) (map[string]*Recording, map[string]*error) {
	errors := make(map[string]*error)
	recs := make(map[string]*Recording)

	// no filters, be safe and don't do anything
	if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
		return recs, errors
	}

	// at least one filter selected
    for id, rec := range GetAllRecordings(nil) {
		roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
		meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
			if rec, err := DeleteRecording(rec.MeetingId); err != nil {
				errors[id] = &err
			} else {
				recs[id] = rec
			}
		}
    }

	return recs, errors
}

func UpdateAllRecordings(filters *RecordingFilters, params *Recording) (map[string]*Recording, map[string]*error) {
	errors := make(map[string]*error)
	recs := make(map[string]*Recording)

	// no filters, be safe and don't do anything
	if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
		return recs, errors
	}

	// at least one filter selected
    for id, rec := range GetAllRecordings(nil) {
		roomMatches := len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds)
		meetingMatches := len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds)
        if roomMatches || meetingMatches {
			if rec, err := UpdateRecording(rec.MeetingId, params); err != nil {
				errors[id] = &err
			} else {
				recs[id] = rec
			}
		}
    }

	return recs, errors
}

func DeleteRecording(uid string) (rec *Recording, err error) {
	rec = RecordingList[uid]
	delete(RecordingList, uid)
	return rec, nil
}

func UpdateRecording(uid string, params *Recording) (a *Recording, err error) {
	if r, ok := RecordingList[uid]; ok {
		if params.Name != "" {
			r.Name = params.Name
		}
		// if params.Published != nil {
		// 	r.Published = params.Published
		// }
		return r, nil
	}
	return nil, errors.New("Recording does not exist")
}
