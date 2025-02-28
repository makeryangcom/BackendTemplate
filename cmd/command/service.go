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

package command

import (
	"log"

	"github.com/gookit/color"
	"github.com/makeryangcom/backend/config"
	"github.com/makeryangcom/backend/pkg/version"
	"github.com/makeryangcom/backend/service"
	"github.com/spf13/cobra"
)

func Service() *cobra.Command {
	command := &cobra.Command{
		Use:     "service",
		Short:   "Start service module",
		Long:    "Start service module",
		Example: "geekros service",
		Run:     serviceRun,
	}
	return command
}

func serviceRun(cmd *cobra.Command, args []string) {

	log.Println(color.Gray.Text("[command]"), color.Gray.Text("serviceRun"))

	version.Get = version.New()

	config.Get = config.New().LoadConfig()

	service.Get = service.New()
	service.Get.Start(func() {
		log.Println(color.Gray.Text("[command]"), color.Gray.Text("serviceRun"), color.Gray.Text("callback"))
	}, func() {
		log.Println(color.Gray.Text("[command]"), color.Gray.Text("serviceRun"), color.Gray.Text("exit"))
	})
}
