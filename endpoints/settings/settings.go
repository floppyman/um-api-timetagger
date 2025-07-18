package settings

import (
	"encoding/json"
	
	"github.com/floppyman/um-api-timetagger/base"
)

func (re SettingEndPoint) Get() (bool, error, GetResult) {
	req := base.CreateRequest(base.HttpGet, "/settings", nil)
	
	ok1, body, err1 := base.DoRequest(req)
	if !ok1 {
		return false, err1, GetResult{}
	}
	
	var res GetResult
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, GetResult{}
	}
	
	return true, nil, res
}

func (re SettingEndPoint) Put(objects []SettingObject) (bool, error, PutResult) {
	bodyBytes, err1 := json.Marshal(objects)
	if err1 != nil {
		return false, err1, PutResult{}
	}
	
	req := base.CreateRequest(base.HttpPut, "/settings", bodyBytes)
	
	ok1, body, err2 := base.DoRequest(req)
	if !ok1 {
		return false, err2, PutResult{}
	}
	
	var res PutResult
	ok2, err3 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err3, PutResult{}
	}
	
	return true, nil, res
}
