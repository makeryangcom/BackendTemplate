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

package main

import (
	"fmt"
	"os"

	"github.com/makeryangcom/backend/cmd/command"
	"github.com/makeryangcom/backend/pkg/version"
	"github.com/spf13/cobra"
)

func main() {

	Version := version.New()

	cmd := &cobra.Command{
		Use:   Version.Name,
		Short: Version.Describe,
		Long:  fmt.Sprintf("%s - %s %s (%s)", Version.Name, Version.Describe, Version.Version, Version.Site),
	}

	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.AddCommand(command.Version())

	cmd.AddCommand(command.Service())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
