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

package version

var Get = &Version{}

type Version struct {
	Name     string
	Describe string
	Site     string
	Version  string
}

func New() *Version {
	return &Version{
		Name:     "backend",
		Describe: "Backend Template",
		Site:     "https://github.com/makeryangcom",
		Version:  "0.0.1",
	}
}
