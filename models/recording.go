package models

import (
	_ "errors"
	"strconv"
	"time"
)

var (
	RecordingList map[string]*Recording
)

func init() {
	RecordingList = make(map[string]*Recording)
	r := Recording{"recording_1", "A Nice Recording"}
	RecordingList["recording_1"] = &r
	r = Recording{"recording_2", "Another Nice Recording"}
	RecordingList["recording_2"] = &r
}

type Recording struct {
	Id   string
	Name string
}

func AddRecording(r Recording) string {
	r.Id = "recording_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	RecordingList[r.Id] = &r
	return r.Id
}

func GetAllRecordings() map[string]*Recording {
	return RecordingList
}
