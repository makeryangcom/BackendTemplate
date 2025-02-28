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

package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/makeryangcom/backend/pkg/utils"
	"gopkg.in/yaml.v3"
)

var Get = &Config{}

type Config struct {
	Path      string   `yaml:"-"`
	Workspace string   `yaml:"-"`
	Runtime   string   `yaml:"-"`
	Server    service  `yaml:"server"`
	Database  database `yaml:"database"`
}

type service struct {
	Mode         string        `yaml:"mode"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func New() *Config {

	workspace := "/opt/geekros"

	if utils.CheckDevMode() {
		workspace = ".."
	}

	configPath := filepath.Join(workspace, "/release/config.sample.yaml")

	return &Config{
		Path:      configPath,
		Workspace: workspace,
		Runtime:   filepath.Join(workspace, "/runtime"),
	}
}

func (config *Config) LoadConfig() *Config {

	if _, err := os.Stat(config.Path); os.IsNotExist(err) {
		config.Server.Mode = "debug"
	}

	file, err := os.ReadFile(config.Path)
	if err != nil {
		return config
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return config
	}

	return config
}

func (c *Config) UpdateConfig() error {

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.Path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
