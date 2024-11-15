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
	"github.com/backend/template/package/version"
	"github.com/backend/template/utils"
	"github.com/gin-gonic/gin"
)

type responseIndex struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Describe string `json:"describe"`
}

func Index(c *gin.Context) {

	returnData := responseIndex{}

	returnData.Name = version.Get.Name
	returnData.Version = version.Get.Version
	returnData.Describe = version.Get.Describe

	utils.Success(c, returnData)
}
