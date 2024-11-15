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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/backend/template/package/config"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

type logData struct {
	Scheme   string `json:"scheme"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	HostName string `json:"host_name"`
	Length   string `json:"length"`
}

func requestLog(c *gin.Context, status string, response string) {

	data := &logData{}

	scheme := "http://"

	if proto := c.Request.Header.Get("x-forwarded-proto"); proto != "" {
		if proto == "https" {
			scheme = "https://"
		}
	}

	if c.Request.TLS != nil {
		scheme = "https://"
	}

	if c.Request.Header.Get("upgrade") == "websocket" {
		if scheme == "https://" {
			scheme = "wss://"
		} else {
			scheme = "ws://"
		}
	}

	data.Scheme = scheme
	data.Domain = c.Request.Host
	data.Path = c.Request.URL.Path
	data.HostName, _ = os.Hostname()

	clientTimeStr := c.Request.Header.Get("robot-x-time")
	if clientTimeStr == "" {
		clientTimeStr = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	}
	clientTime, err := strconv.ParseInt(clientTimeStr, 10, 64)
	if err != nil {
		data.Length = fmt.Sprintf("%.2fms", 0.00)
	}
	duration, err := CalculateRequestDuration(clientTime)
	if err != nil {
		data.Length = fmt.Sprintf("%.2fms", 0.00)
	} else {
		data.Length = fmt.Sprintf("%.2fms", duration)
	}

	dataString, _ := json.Marshal(data)

	var request string
	if c.Request.Method == "GET" {
		request = c.Request.URL.Query().Encode()
	}
	if c.Request.Method == "POST" {
		request = c.Request.Form.Encode()
		if len(request) > 100 && config.Get.Server.Mode == "release" {
			midpoint := len(request) / 2
			request = request[:midpoint-50] + "..." + request[midpoint+50:]
		}
	}

	icon := color.Green.Text(fmt.Sprintf("%-19s⇐", ""))
	if status == "log" {
		icon = color.Gray.Text(fmt.Sprintf("%-19s⇐", ""))
	}
	if status == "warning" {
		icon = color.Yellow.Text(fmt.Sprintf("%-19s⇐", ""))
	}
	if status == "error" {
		icon = color.Red.Text(fmt.Sprintf("%-19s⇐", ""))
	}

	if len(response) > 100 && config.Get.Server.Mode == "release" {
		midpoint := len(response) / 2
		response = response[:midpoint-50] + "..." + response[midpoint+50:]
	}

	if status == "log" {
		c.Request.Method = color.Gray.Text(c.Request.Method)
		request = color.Gray.Text(request)
		response = color.Gray.Text(response)
	}

	log.Println(
		string(dataString), "\n",
		color.Gray.Text(fmt.Sprintf("%-19s⇒", "")), c.Request.Method, request, "\n",
		color.Gray.Text(fmt.Sprintf("%-19s⇒", "")), fmt.Sprintf("%s -> %s", time.Unix(0, clientTime*int64(time.Millisecond)).Format(time.RFC3339), time.Now().Format(time.RFC3339)), "\n",
		color.Gray.Text(fmt.Sprintf("%-19s⇒", "")), c.Request.Header.Get("robot-x-source"), c.ClientIP(), c.Request.Header.Get("robot-x-ip"), c.Request.Header.Get("robot-x-device"), "\n",
		icon,
		response)
}

func Success(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 0, "message": "success", "data": data})

	requestLog(c, "success", string(logJson))
}

func Error(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": "error",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 10000, "message": "error", "data": data})

	requestLog(c, "error", string(logJson))
}

func Warning(c *gin.Context, code int, message string, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": code, "message": message, "data": data})

	requestLog(c, "warning", string(logJson))
}

func Log(c *gin.Context, message string, data interface{}) {

	logJson, _ := json.Marshal(gin.H{"code": 90000, "message": message, "data": data})

	requestLog(c, "log", string(logJson))
}
