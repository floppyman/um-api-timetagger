package updates

import (
	"github.com/floppyman/um-api-timetagger/endpoints/records"
	"github.com/floppyman/um-api-timetagger/endpoints/settings"
)

type UpdateEndPoint struct {
}

type UpdateObject struct {
	ServerTime float64                  `json:"server_time"`
	Reset      byte                     `json:"reset" default:"0"`
	Records    []records.RecordObject   `json:"records"`
	Settings   []settings.SettingObject `json:"settings"`
}
