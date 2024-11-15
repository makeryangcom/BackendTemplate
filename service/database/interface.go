// Copyright 2024 GEEKROS, Inc.
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
	"time"

	"github.com/jinzhu/gorm"
)

type Interface struct {
	TableName string
}

func Table(table string) *Interface {
	return &Interface{
		TableName: table,
	}
}

func (i *Interface) CreateData(data interface{}) error {
	err := Get.db.Table(i.TableName).Create(data).Error
	return err
}

func (i *Interface) UpdateData(query interface{}, data map[string]interface{}) error {
	data["update_at"] = time.Now().Unix()
	err := Get.db.Table(i.TableName).Where(query).Updates(data).Error
	return err
}

func (i *Interface) ExprData(query interface{}, field string, operation string, data int) error {
	return Get.db.Table(i.TableName).Where(query).Updates(map[string]interface{}{
		field:       gorm.Expr(field+" "+operation+" ?", data),
		"update_at": time.Now().Unix(),
	}).Error
}

func (i *Interface) GetData(dataStruct interface{}, query interface{}, order string) error {
	err := Get.db.Table(i.TableName).Where(query).Order(order).First(dataStruct).Error
	return err
}

func (i *Interface) ListData(dataStruct interface{}, query interface{}, order string, limit int) error {
	err := Get.db.Table(i.TableName).Where(query).Order(order).Limit(limit).Find(dataStruct).Error
	return err
}

func (i *Interface) PageData(dataStruct interface{}, query interface{}, order string, limit int, page int) error {
	if limit > 50 {
		limit = 50
	}
	err := Get.db.Table(i.TableName).Where(query).Order(order).Limit(limit).Offset(page * limit).Find(dataStruct).Error
	return err
}

func (i *Interface) CountData(query interface{}) (int, error) {
	count := 0
	err := Get.db.Table(i.TableName).Where(query).Count(&count).Error
	return count, err
}

func (i *Interface) DeleteData(dataStruct interface{}, query interface{}) error {
	err := Get.db.Table(i.TableName).Where(query).Delete(dataStruct).Error
	return err
}
