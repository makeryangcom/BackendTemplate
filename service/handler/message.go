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

package handler

import (
	"encoding/json"

	"github.com/backend/template/package/socket"
	"github.com/backend/template/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func MessageIndex(c *gin.Context) {

	if !socket.Get.IsUpgrade(c.Request) {
		utils.Warning(c, 10000, "Please use ws or wss protocol", utils.EmptyData{})
		return
	}

	conn, err := socket.Get.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.Warning(c, 10000, "Failed to upgrade to socket", utils.EmptyData{})
		return
	}
	defer conn.Close()

	if !socket.Get.Status {
		socket.Get.User = make(map[*websocket.Conn]bool)
		socket.Get.Status = true
	}

	socket.Get.User[conn] = true

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message socket.MessageFormat
		err = json.Unmarshal(data, &message)
		if err == nil {
			if message.Command != "" {
				socket.Get.SendMessage(message.Command, message.Message, message.Data)
			}
		}
	}

	utils.Log(c, "socket connection", utils.EmptyData{})
}
