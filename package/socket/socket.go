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

package socket

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Get = &Socket{}

type Socket struct {
	Status  bool
	Upgrade websocket.Upgrader
	User    map[*websocket.Conn]bool
	mutex   sync.Mutex
}

type MessageFormat struct {
	Command string      `json:"command"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New() *Socket {
	return &Socket{
		Upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		User: map[*websocket.Conn]bool{},
	}
}

func (s *Socket) IsUpgrade(request *http.Request) bool {
	return websocket.IsWebSocketUpgrade(request)
}

func (s *Socket) SendMessage(command string, message string, data interface{}) {
	send := MessageFormat{}
	send.Command = command
	send.Message = message
	send.Data = data
	messageJson, _ := json.Marshal(send)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for user := range s.User {
		err := user.WriteMessage(websocket.TextMessage, messageJson)
		if err != nil {
			user.Close()
			delete(s.User, user)
		}
	}
}
