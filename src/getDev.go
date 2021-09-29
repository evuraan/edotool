/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	devices  = "/proc/bus/input/devices"
	touchpad = "TouchPad"
	keyboard = "keyboard"
)

var (
	touchPadIdentifier = ""
)

func getDeviceForPattern(lookForA string) (devicePath string) {
	if len(lookForA) < 1 {
		return
	}
	lookFor := strings.ToLower(lookForA)
	fmap := parseFileToMap(devices)
	if len(fmap) < 1 {
		return
	}
	var entry, lentr string
	pat := regexp.MustCompile(`event\d*`)
	for k := range fmap {
		entry = fmap[k]
		if len(entry) < 1 {
			continue
		}
		lentr = strings.ToLower(entry)
		if strings.Contains(lentr, lookFor) {
			pos := k + 4
			if len(fmap) < pos {
				continue
			}
			handler := fmap[pos]
			s := pat.FindString(handler)

			if len(s) > 0 {
				devicePath = fmt.Sprintf("/dev/input/%s", s)
				return devicePath
			}
		}
	}

	return devicePath
}

func parseFileToMap(configFile string) (someDict map[int]string) {
	someDict = make(map[int]string)

	func() {
		f, err := os.Open(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open config file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		fscanner := bufio.NewScanner(f)
		for fscanner.Scan() {
			line := fscanner.Text()
			if len(line) > 0 {
				someDict[len(someDict)] = line
			}
		}
	}()

	return someDict
}

func showDevices() {
	fmap := parseFileToMap(devices)
	if len(fmap) < 1 {
		return
	}
	devToName := make(map[string]string)
	pat := regexp.MustCompile(`event\d*`)

	for i := range fmap {
		someLine := fmap[i]
		if strings.HasPrefix(someLine, "N: Name=") {
			name := someLine[8:]
			if len(name) > 1 {
				name = strings.ReplaceAll(name, "\"", "")
				if len(name) <= 1 {
					continue
				}
				pos := i + 4
				if len(fmap) < pos {
					continue
				}
				eventLine := fmap[pos]
				s := pat.FindString(eventLine)

				if len(s) > 0 {
					devicePath := fmt.Sprintf("/dev/input/%s", s)
					devToName[devicePath] = name
				}
			}
		}
	}

	if len(devToName) < 1 {
		fmt.Fprintf(os.Stderr, "Error: Could not find any input devices.\n")
		os.Exit(1)
	} else {
		fmt.Printf("Available: %d devices\n", len(devToName))
	}

	for i := 0; i < len(devToName); i++ {
		key := fmt.Sprintf("/dev/input/event%d", i)
		rhs, ok := devToName[key]
		if ok {
			fmt.Printf("%s:\t%s\n", key, rhs)
		}
	}

}
