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

package database

import (
	"fmt"
	"log"
	"time"

	"github.com/gookit/color"
	"github.com/jinzhu/gorm"
	"github.com/makeryangcom/backend/config"
)

var Get = &Database{}

type Database struct {
	db           *gorm.DB
	DefaultField DefaultField
}

type DefaultField struct {
	CreateAt int `gorm:"Column:create_at" json:"create_at"`
	UpdateAt int `gorm:"Column:update_at" json:"update_at"`
	DeleteAt int `gorm:"Column:delete_at" json:"delete_at"`
}

func New() *Database {
	return &Database{}
}

func (d *Database) Init() *Database {

	var err error

	d.db, err = gorm.Open(
		config.Get.Database.Type,
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Get.Database.User,
			config.Get.Database.Password,
			config.Get.Database.Host,
			config.Get.Database.Name,
		),
	)
	if err != nil {
		log.Println(color.Red.Text(err.Error()))
	}

	if config.Get.Server.Mode == "release" {
		d.db.LogMode(false)
	} else {
		d.db.LogMode(true)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	d.db.SingularTable(true)

	d.db.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if !scope.HasError() {
			nowTime := time.Now().Unix()
			if createTimeField, ok := scope.FieldByName("CreateAt"); ok {
				if createTimeField.IsBlank {
					err := createTimeField.Set(nowTime)
					if err != nil {
						return
					}
				}
			}

			if modifyTimeField, ok := scope.FieldByName("UpdateAt"); ok {
				if modifyTimeField.IsBlank {
					err := modifyTimeField.Set(nowTime)
					if err != nil {
						return
					}
				}
			}
		}
	})

	d.db.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.FieldByName("UpdateAt"); ok {
			err := scope.SetColumn("UpdateAt", time.Now().Unix())
			if err != nil {
				return
			}
		}
	})

	d.db.DB().SetMaxIdleConns(1000)
	d.db.DB().SetMaxOpenConns(10000)

	d.db.DB().SetConnMaxLifetime(time.Second * 45)

	return d
}
