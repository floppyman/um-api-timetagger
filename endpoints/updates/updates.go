package updates

import (
	"fmt"
	
	"github.com/floppyman/um-api-timetagger/base"
)

func (re UpdateEndPoint) Get(since int64) (bool, error, UpdateObject) {
	req := base.CreateRequest(base.HttpGet, fmt.Sprintf("/updates?since=%d", since), nil)
	
	ok1, body, err1 := base.DoRequest(req)
	if !ok1 {
		return false, err1, UpdateObject{}
	}
	
	var res UpdateObject
	ok2, err2 := base.UnpackBody(body, &res)
	if !ok2 {
		return false, err2, UpdateObject{}
	}
	
	return true, nil, res
}
