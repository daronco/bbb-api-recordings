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
        if len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds) {
			ret[id] = rec
		} else if len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds) {
			ret[id] = rec
		}
    }
	return ret
}

func DeleteAllRecordings(filters *RecordingFilters) map[string]*error {
	ret := make(map[string]*error)

	// no filters, be safe and don't do anything
	if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
		return ret
	}

	// at least one filter selected
    for id, rec := range GetAllRecordings(nil) {
        if len(filters.RoomIds) > 0 && utils.StringInSlice(rec.RoomId, filters.RoomIds) {
			if _, err := DeleteRecording(rec); err != nil {
				ret[id] = &err
			}
		} else if len(filters.MeetingIds) > 0 && utils.StringInSlice(rec.MeetingId, filters.MeetingIds) {
			if _, err := DeleteRecording(rec); err != nil {
				ret[id] = &err
			}
		}
    }

	return ret
}

func DeleteRecording(rec *Recording) (b bool, err error) {
	delete(RecordingList, rec.MeetingId)
	return true, nil
}

func UpdateAllRecordings(params *Recording) (a bool, err error) {
	for _, r := range RecordingList {
		_, lerr := UpdateRecording(r.MeetingId, params)
		if lerr != nil {
			return false, lerr
		}
	}
	return true, nil
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
