package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	RecordingList map[string]*Recording
)

func init() {
	RecordingList = make(map[string]*Recording)
	r1 := Recording{"recording_1", "A Nice Recording", true}
	RecordingList["recording_1"] = &r1
	r2 := Recording{"recording_2", "Another Nice Recording", false}
	RecordingList["recording_2"] = &r2
}

type Recording struct {
	Id        string
	Name      string
	Published bool
}

func AddRecording(r Recording) string {
	r.Id = "recording_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	RecordingList[r.Id] = &r
	return r.Id
}

func GetAllRecordings() map[string]*Recording {
	return RecordingList
}

func DeleteAllRecordings() (a bool, err error) {
	for _, r := range RecordingList {
		DeleteRecording(r.Id)
	}
	return true, nil
}

func UpdateAllRecordings(params *Recording) (a bool, err error) {
	for _, r := range RecordingList {
		_, lerr := UpdateRecording(r.Id, params)
		if lerr != nil {
			return false, lerr
		}
	}
	return true, nil
}

func DeleteRecording(uid string) {
	delete(RecordingList, uid)
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
