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
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetLocalIPAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Check if the interface is up and running
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagRunning == 0 {
			continue // Ignore interfaces that are not up and running
		}

		// Optionally check if the interface has transmitted or received any packets
		if iface.Flags&net.FlagLoopback != 0 {
			continue // Ignore loopback interfaces
		}

		addrs, addrErr := iface.Addrs()
		if addrErr != nil {
			continue
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if ipv4 := v.IP.To4(); ipv4 != nil {
					return ipv4.String(), nil // Return the first active IPv4 address found
				}
			}
		}
	}

	return "", fmt.Errorf("no active network interface found")
}

func GetWifiSerialNumber() (string, error) {
	cmd := exec.Command("lshw", "-C", "network")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	re := regexp.MustCompile(`(?s)\*-wifi:0(?:.|\n)*?serial: ([\da-f:]+)`)
	match := re.FindStringSubmatch(out.String())
	if len(match) < 2 {
		return "", errors.New("not serial number")
	}

	return match[1], nil
}

func GetSocUid(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("command does not exist at path: %s", path)
		}
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("path points to a directory, not a command: %s", path)
	}

	cmd := exec.Command(path)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	outputStr := string(output)
	parts := strings.Split(outputStr, "soc_uid: ")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1]), nil
	}
	return "", fmt.Errorf("soc_uid not found in the output")
}

func IsGraphicalTargetActive() bool {
	cmd := exec.Command("systemctl", "is-active", "graphical.target")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return string(output) == "active\n"
}
