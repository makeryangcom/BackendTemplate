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

package crontab

import (
	"github.com/robfig/cron/v3"
)

var Get = &Crontab{}

type Crontab struct {
	cron *cron.Cron
}

func New() *Crontab {
	return &Crontab{}
}

func (c *Crontab) Start() *Crontab {
	c.cron = cron.New(cron.WithSeconds())
	_, _ = c.cron.AddFunc("*/1 * * * * *", func() {
		// log.Println(color.White.Text(fmt.Sprintf("%s", "crontab 1 second")))
	})
	_, _ = c.cron.AddFunc("*/10 * * * * *", func() {
		// log.Println(color.White.Text(fmt.Sprintf("%s", "crontab 10 second")))
	})
	_, _ = c.cron.AddFunc("0 */1 * * * *", func() {
		// log.Println(color.Green.Text("crontab 1 minute"))
	})

	c.cron.Start()

	return c
}
