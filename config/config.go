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
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var Get = &Config{}

type Config struct {
	Path      string   `yaml:"-"`
	Workspace string   `yaml:"-"`
	Runtime   string   `yaml:"-"`
	Server    service  `yaml:"server"`
	Frontend  frontend `yaml:"frontend"`
	Database  database `yaml:"database"`
	Hash      hash     `yaml:"hash"`
}

type frontend struct {
	Path string `yaml:"path"`
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

type hash struct {
	Salt     string `yaml:"salt"`
	Alphabet string `yaml:"alphabet"`
}

func New() *Config {

	workspace := "/opt/backend"

	if CheckDevMode() {
		workspace, _ = os.Getwd()
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
		config.Server.Port = 8090
		config.Server.ReadTimeout = 60 * time.Second
		config.Server.WriteTimeout = 60 * time.Second

		config.Frontend.Path = "/opt/frontend"

		config.Database.Type = "mysql"
		config.Database.Host = "127.0.0.1"
		config.Database.Name = "example"
		config.Database.User = "root"
		config.Database.Password = "example"

		config.Hash.Salt = "example"
		config.Hash.Alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz123456789"

		config.UpdateConfig()
	}

	file, err := os.ReadFile(config.Path)
	if err != nil {
		return config
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return config
	}

	if config.Frontend.Path == "" {
		config.Frontend.Path = "/opt/frontend"
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

func CheckDevMode() bool {
	mode := false
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		mode = true
	}
	return mode
}
