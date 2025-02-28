// Copyright 2024 MakerYang, Inc.
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

package router

import (
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/makeryangcom/backend/config"
	"github.com/makeryangcom/backend/service/controller"
)

type Router struct {
	engine *gin.Engine
}

func New() *Router {

	return &Router{}
}

func (router *Router) Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
	}
}

func (router *Router) Initialization(mode string) *Router {

	gin.SetMode(mode)

	router.engine = gin.New()
	router.engine.Use(gin.Recovery())
	router.engine.Use(cors.Default())
	router.engine.Use(router.Authentication())

	return router
}

func (router *Router) Handler() *gin.Engine {
	router.FrontendHandler()

	log.Println(color.Gray.Text("[router]"), color.Gray.Text("handler"))

	router.engine.GET("/index", controller.Index)

	return router.engine
}

func (router *Router) FrontendHandler() *gin.Engine {

	log.Println(color.Gray.Text("[router]"), color.Gray.Text("frontend handler"))

	basePath := filepath.Join(config.Get.Frontend.Path, "/release")

	staticDirs := []string{"assets", "images", "locales", "videos"}
	for _, dir := range staticDirs {
		router.engine.Static("/"+dir, filepath.Join(basePath, dir))
	}

	router.engine.GET("/icon.svg", func(c *gin.Context) {
		c.File(filepath.Join(basePath, "/release/icon.svg"))
	})

	router.engine.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(basePath, "/release/index.html"))
	})

	return router.engine
}
