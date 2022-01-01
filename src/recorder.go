/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	maxRecordTime = 300
)

func doRecord() {
	fmt.Printf("Recoding device: %s\n", inputDevice)
	recordFn(inputDevice, 10)
}

func recordFn(device string, duration int) bool {

	if duration < 1 {
		fmt.Fprintf(os.Stderr, "Invalid record interval\n")
		os.Exit(1)
	}

	if duration > maxRecordTime {
		fmt.Fprintf(os.Stderr, "recording span %d exceeds max limit. exiting\n", duration)
		os.Exit(1) // hard exit
	}

	outputFile := fmt.Sprintf("%s/edotool-Recorded-%d.skit", os.TempDir(), time.Now().Unix())

	done := make(chan bool, 1)
	defer close(done)

	mu := &sync.RWMutex{}
	keepRecording := true
	tempMap := make(map[int]string)

	go func() {
		fPtr, err := os.Open(device)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Err :%v\n", err)
			return
		}
		defer fPtr.Close()
		defer func() {
			done <- true
		}()
		data := make([]byte, 24)
		for true {
			n, err := fPtr.Read(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Err :%v\n", err)
				break
			}
			if n < 24 {
				continue
			}
			typeThing := binary.LittleEndian.Uint16(data[16:18])
			codeThing := binary.LittleEndian.Uint16(data[18:20])
			value := binary.LittleEndian.Uint32(data[20:24])
			appendThis := fmt.Sprintf("%d|%d|%d\n", typeThing, codeThing, value)
			tempMap[len(tempMap)] = appendThis
			mu.Lock()
			x := keepRecording
			mu.Unlock()
			if !x {
				break
			}
		}
	}()
	fmt.Printf("Recording from %s to file %s, duration %d seconds\n", device, outputFile, duration)
	xo := 0
	tmo := time.Duration(duration * int(time.Second))
	select {
	case <-done:
		goto recordCleanup
	case <-time.After(tmo):
		xo++
		goto recordCleanup
	}

recordCleanup:
	mu.Lock()
	keepRecording = false
	mu.Unlock()
	if len(tempMap) < 1 {
		fmt.Fprintf(os.Stderr, "Did not record any events\n")
		return false
	} else {
		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "recordFn err %v\n", err)
			os.Exit(1) // hard exit
		}
		defer f.Close()
		hvst := 0
		for i := 0; i < len(tempMap); i++ {
			writeThis, ok := tempMap[i]
			if !ok || len(writeThis) < 1 {
				continue
			}
			f.WriteString(writeThis)
			hvst++
		}
		fmt.Printf("\nDone!\nHarvested %d events to %s\n", hvst, outputFile)
	}
	return xo > 0
}
