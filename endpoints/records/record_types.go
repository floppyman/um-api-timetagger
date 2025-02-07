package records

type RecordEndPoint struct {
}

type GetResult struct {
	Records []RecordObject `json:"records"`
}

type RecordObject struct {
	Key          string  `json:"key"`
	StartTime    int64   `json:"t1"`
	StopTime     int64   `json:"t2"`
	Description  string  `json:"ds"`
	ModifiedTime int64   `json:"mt"`
	ServerTime   float64 `json:"st" default:"0.0"`
}

type PutResult struct {
	Accepted []string `json:"accepted"`
	Failed   []string `json:"failed"`
	Errors   []string `json:"errors"`
}
