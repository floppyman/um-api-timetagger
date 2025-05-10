package timetagger

import (
	"github.com/spf13/viper"
	
	"github.com/floppyman/um-api-timetagger/base"
	"github.com/floppyman/um-api-timetagger/endpoints/records"
	"github.com/floppyman/um-api-timetagger/endpoints/settings"
	"github.com/floppyman/um-api-timetagger/endpoints/updates"
)

//goland:noinspection GoUnusedExportedFunction
func NewClient(url string, token string) TimeTaggerApiClient {
	return newClient(base.TimeTaggerOptions{
		ApiUrl:   url,
		ApiToken: token,
	})
}

func NewClientFromOptions(conf *viper.Viper) TimeTaggerApiClient {
	return newClient(base.TimeTaggerOptions{
		ApiUrl:   conf.GetString("api_url"),
		ApiToken: conf.GetString("api_token"),
	})
}

func newClient(options base.TimeTaggerOptions) TimeTaggerApiClient {
	base.Init(options)
	return TimeTaggerApiClient{
		Records:  records.RecordEndPoint{},
		Settings: settings.SettingEndPoint{},
		Updates:  updates.UpdateEndPoint{},
	}
}

//goland:noinspection GoNameStartsWithPackageName
type TimeTaggerApiClient struct {
	Records  records.RecordEndPoint
	Settings settings.SettingEndPoint
	Updates  updates.UpdateEndPoint
}
