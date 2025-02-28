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

package utils

import (
	"fmt"
	"time"
)

func CalculateRequestDuration(clientTime int64) (float64, error) {
	clientTimeObj := time.Unix(0, clientTime*int64(time.Millisecond))

	now := time.Now()

	duration := now.Sub(clientTimeObj)

	if duration < 0 {
		return 0, fmt.Errorf("client time is later than server time, clocks may be out of sync")
	}

	durationMs := float64(duration.Nanoseconds()) / float64(time.Millisecond)
	roundedDurationMs := float64(int(durationMs*100)) / 100

	return roundedDurationMs, nil
}
