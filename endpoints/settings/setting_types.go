package settings

type SettingEndPoint struct {
}

type GetResult struct {
	Settings []SettingObject `json:"settings"`
}

type SettingObject struct {
	Key          string      `json:"key"`
	Value        interface{} `json:"value"`
	ModifiedTime int64       `json:"mt"`
	ServerTime   float64     `json:"st" default:"0.0"`
}

type PutResult struct {
	Accepted []string `json:"accepted"`
	Failed   []string `json:"failed"`
	Errors   []string `json:"errors"`
}
