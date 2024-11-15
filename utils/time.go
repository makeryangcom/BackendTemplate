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

package utils

import (
	"fmt"
	"strconv"
	"time"
)

func TimeFormat(unix int) string {
	timeInt := time.Unix(int64(unix), 0)
	return timeInt.Format("2006-01-02 15:04:05")
}

func TimeMinFormat(unix int) string {
	timeInt := time.Unix(int64(unix), 0)
	return timeInt.Format("2006-01-02")
}

func DateFormat(times int) string {
	createTime := time.Unix(int64(times), 0)
	now := time.Now().Unix()

	difTime := now - int64(times)

	str := ""
	if difTime < 60 {
		str = "just now"
	} else if difTime < 3600 {
		M := difTime / 60
		str = strconv.Itoa(int(M)) + " minutes ago"
	} else if difTime < 3600*24 {
		H := difTime / 3600
		str = strconv.Itoa(int(H)) + " hours ago"
	} else {
		str = createTime.Format("2006-01-02 15:04:05")
	}

	return str
}

func DateFormatFine(timestamp int64) string {
	now := time.Now()
	past := time.Unix(timestamp, 0)
	duration := now.Sub(past)

	seconds := int(duration.Seconds())
	minutes := int(duration.Minutes())
	hours := int(duration.Hours())
	days := int(duration.Hours() / 24)
	weeks := days / 7
	months := int(now.Sub(past).Hours() / 24 / 30)

	switch {
	case seconds < 60:
		return fmt.Sprintf("%d seconds ago", seconds)
	case minutes < 60:
		return fmt.Sprintf("%d minutes ago", minutes)
	case hours < 24:
		return fmt.Sprintf("%d hours ago", hours)
	case days < 7:
		return fmt.Sprintf("%d days ago", days)
	case days < 30:
		return fmt.Sprintf("%d weeks ago", weeks)
	case months < 12:
		return fmt.Sprintf("%d months ago", months)
	default:
		return past.Format("2006-01-02 15:04:05")
	}
}

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
