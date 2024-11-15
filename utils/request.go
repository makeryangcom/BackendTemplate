// Copyright 2024 ARMCNC, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

var CloudServiceHost = "https://backend.geekros.com"

type ResponseService struct {
	Code    int     `json:"code"`
	Data    service `json:"data"`
	Message string  `json:"message"`
}

type service struct {
	Sign  string `json:"sign"`
	Token string `json:"token"`
}

func Service(path string, method string, parameters map[string]string, data map[string]string) (*http.Response, ResponseService, error) {
	responseData := ResponseService{}
	responseData.Code = 10000

	var bodyData []byte
	if data != nil {
		bodyData, _ = json.Marshal(data)
	}

	request, err := http.NewRequest(method, CloudServiceHost+path, bytes.NewReader(bodyData))
	if err != nil {
		return nil, responseData, err
	}

	request.Header.Set("Content-type", "application/json")

	socuid, err := GetSocUid("/usr/bin/hrut_socuid")
	if err != nil {
		return nil, responseData, err
	}

	request.Header.Set("Content-X-Device", socuid)

	request.Header.Set("Content-X-Referer", "")

	request.Header.Set("Content-X-Source", "terminal")

	ip, err := GetLocalIPAddress()
	if err != nil {
		return nil, responseData, err
	}

	request.Header.Set("Content-X-IP", ip)

	request.Header.Set("Content-X-Time", strconv.FormatInt(time.Now().UnixMilli(), 10))

	query := request.URL.Query()
	if parameters != nil {
		for key, val := range parameters {
			query.Add(key, val)
		}
		request.URL.RawQuery = query.Encode()
	}

	client := &http.Client{}

	response, _ := client.Do(request)

	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &responseData)
	return response, responseData, err
}
