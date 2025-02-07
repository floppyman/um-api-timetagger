package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/umbrella-sh/um-common/logging/ulog"
)

var Options TimeTaggerOptions

type HttpMethod string

const (
	HttpGet = "GET"
	HttpPut = "PUT"
)

func Init(options TimeTaggerOptions) {
	Options = options
}

func CreateRequest(method HttpMethod, urlPath string, body []byte) *http.Request {
	fullUrl := fmt.Sprintf("%s%s", Options.ApiUrl, urlPath)
	ulog.Console.Debug().Msgf("TimeTagger Url: %s", fullUrl)

	var req *http.Request
	var err error

	if method != HttpGet && body != nil && len(body) > 0 {
		req, err = http.NewRequest(string(method), fullUrl, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(string(method), fullUrl, nil)
	}

	if err != nil {
		return nil
	}

	req.Header.Add("authtoken", Options.ApiToken)
	if method != HttpGet {
		req.Header.Add("Content-Type", "application/json")
	}

	return req
}

func DoRequest(req *http.Request) (bool, []byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}

	return true, body, nil
}

func UnpackBody(body []byte, res any) (bool, error) {
	err := json.Unmarshal(body, &res)
	if err != nil {
		return false, err
	}

	return true, nil
}
