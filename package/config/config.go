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

package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

var Get = &Config{}

type Config struct {
	Path     string   `yaml:"path"`
	Server   service  `yaml:"server"`
	Database database `yaml:"database"`
	Cdn      cdn      `yaml:"cdn"`
	Hash     hash     `yaml:"hash"`
	Mail     mail     `yaml:"mail"`
	Referer  referer  `yaml:"referer"`
	Auth     auth     `yaml:"auth"`
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

type cdn struct {
	All string `yaml:"all"`
}

type hash struct {
	Salt     string `yaml:"salt"`
	Alphabet string `yaml:"alphabet"`
}

type mail struct {
	Smtp     string `yaml:"smtp"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type referer struct {
	Domain []string `yaml:"domain"`
}

type auth struct {
	Secret     string `yaml:"secret"`
	Expiration int    `yaml:"expiration"`
}

func New() *Config {

	configPath := os.Getenv("BACKEND_CONFIG_PATH")
	if configPath == "" {
		configPath = filepath.Join("../release/", "config.sample.yaml")
	}

	return &Config{Path: configPath}
}

func (c *Config) LoadConfig() *Config {

	file, err := os.ReadFile(c.Path)
	if err != nil {
		return c
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return c
	}

	return c
}
